package sorry

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

const (
	PgxErrorCode         = 1000
	JSONErrorCode        = 2000
	ParseErrorCode       = 3000
	DefaultErrorCode     = 4000
	ValidationErroCode   = 5000
	GRPCErrorCode        = 6000
	TimeoutErrorCode     = 7000
	HTTPRequestErrorCode = 8000
)

// Error fit a error type for handling
type Error struct {
	Msg        string   `json:"msg"`
	Code       int      `json:"code"`
	Trace      []string `json:"-"`
	Err        error    `json:"-"`
	StatusCode int      `json:"-"`
}

// Error enable Error type to implements error type
func (e *Error) Error() string {
	return e.Msg
}

// Unwrap return the specific error cause for this error
func (e *Error) Unwrap() error {
	return e.Err
}

type wrappedError interface {
	Unwrap() error
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// fromErr wraps errors to provide user readable messages
func fromErr(rawError error) error {
	var (
		msg                  string
		code, httpStatusCode int
	)

	switch err := rawError.(type) {
	case *json.UnmarshalTypeError:
		msg, code = rawError.Error(), JSONErrorCode

	case validator.ValidationErrors:
		msg, code = rawError.Error(), ValidationErroCode

	case *reflect.ValueError, *strconv.NumError, *time.ParseError:
		msg, code = rawError.Error(), ParseErrorCode

	case *pgconn.PgError:
		msg, code = rawError.Error(), PgxErrorCode

	case *url.Error:
		msg, code = rawError.Error(), HTTPRequestErrorCode

	case *Error:
		rawError, msg, code, httpStatusCode = err, err.Msg, err.Code, err.StatusCode

	case error:
		switch err {
		case sql.ErrNoRows:
			msg, code, httpStatusCode = err.Error(), DefaultErrorCode, http.StatusNotFound

		case io.EOF:
			msg, code = err.Error(), DefaultErrorCode

		case strconv.ErrSyntax:
			msg, code = err.Error(), ParseErrorCode

		default:
			msg, code, httpStatusCode = rawError.Error(), DefaultErrorCode, http.StatusBadRequest
		}
	case nil:
		return nil

	default:
		msg, code, httpStatusCode = rawError.Error(), DefaultErrorCode, http.StatusBadRequest
	}

	return &Error{
		Msg:        msg,
		Err:        rawError,
		Code:       code,
		StatusCode: httpStatusCode,
	}
}

// Err builds annotated error instance from any error value
func Err(err error) error {
	var e *Error

	if !errors.As(err, &e) || errors.Is(e, err) {
		err = fromErr(err)
	}

	return errors.WithStack(err)
}

// Wrap wraps an error adding an information message
func Wrap(err error, message string) error {
	return errors.Wrap(Err(err), message)
}

// NewErr creates an annotated error instance with default values
func NewErr(message string, statusCode ...int) error {
	sc := http.StatusBadRequest
	if len(statusCode) > 0 {
		sc = statusCode[0]
	}

	return Err(&Error{
		Msg:        message,
		Err:        errors.Errorf("Inline error message: '%s'. See the stack trace of the error for additional information.", message),
		Code:       DefaultErrorCode,
		StatusCode: sc,
	})
}

// Handling handles an error by setting a message and a response status code
func Handling(ctx *gin.Context, err error) {
	var e *Error

	if !errors.As(err, &e) {
		Handling(ctx, Err(err))
		return
	}

	e.Msg = err.Error()
	e.Trace, _ = ReconstructStackTrace(err, e)

	if reqIDVal, hasRID := ctx.Get("RID"); hasRID {
		if reqID, ok := reqIDVal.(string); ok {
			if len(reqID) > 0 {
				e.Msg = fmt.Sprintf("%s [%s]", e.Msg, reqID[:min(6, len(reqID))])
			}
		}
	}

	ctx.JSON(e.StatusCode, e)
	ctx.Set("error", err)
	ctx.Abort()
}

// ReconstructStackTrace tries to reconstruct the stack trace of an error
func ReconstructStackTrace(err error, bound error) ([]string, bool) {
	var (
		wrapped wrappedError
		tracer  stackTracer
		output  []string
		traced  bool
	)

	if errors.As(err, &wrapped) {
		internal := wrapped.Unwrap()

		if !errors.Is(internal, bound) {
			output, traced = ReconstructStackTrace(internal, bound)
		}

		if !traced && errors.As(err, &tracer) {
			stack := tracer.StackTrace()
			for _, frame := range stack {
				output = append(output, fmt.Sprintf("%+v", frame))
			}
			traced = true
		}
	}

	return output, traced
}
