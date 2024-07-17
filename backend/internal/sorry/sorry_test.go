package sorry

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromErr(t *testing.T) {
	tests := []struct {
		name           string
		inputError     error
		expectedCode   int
		expectedStatus int
	}{
		{
			name:           "JSONUnmarshalTypeError",
			inputError:     &json.UnmarshalTypeError{Field: "field", Type: reflect.TypeOf(""), Value: "value"},
			expectedCode:   JSONErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "ValidationErrors",
			inputError:     validator.ValidationErrors{},
			expectedCode:   ValidationErroCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "PgError",
			inputError:     &pgconn.PgError{Severity: "Error", Message: "Error", Code: "23505"},
			expectedCode:   PgxErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "DefaultError",
			inputError:     errors.New("some error"),
			expectedCode:   DefaultErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "URLError",
			inputError:     &url.Error{Op: "op", URL: "url", Err: errors.New("some error")},
			expectedCode:   HTTPRequestErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "ValueError",
			inputError:     &reflect.ValueError{Method: "method", Kind: reflect.Invalid},
			expectedCode:   ParseErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "NumError",
			inputError:     &strconv.NumError{Func: "func", Num: "num", Err: errors.New("some error")},
			expectedCode:   ParseErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "TimeParseError",
			inputError:     &time.ParseError{Layout: "layout", Value: "value", LayoutElem: "layoutElem", ValueElem: "valueElem"},
			expectedCode:   ParseErrorCode,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "ValidationError",
			inputError:     validator.ValidationErrors{},
			expectedCode:   ValidationErroCode,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fromErr(tt.inputError)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}

			e, ok := err.(*Error)
			require.True(t, ok, "Expected error of type *Error, got %T", err)
			assert.Equal(t, tt.expectedCode, e.Code)
			assert.Equal(t, tt.expectedStatus, e.StatusCode)
		})
	}
}

func TestErr(t *testing.T) {
	rawErr := errors.New("raw error message")
	err := Err(rawErr)
	assert.NotNil(t, err)
	assert.Equal(t, rawErr.Error(), err.Error())
}

func TestWrap(t *testing.T) {
	rawErr := errors.New("raw error message")
	wrappedErr := Wrap(rawErr, "additional context")
	assert.NotNil(t, wrappedErr)
	assert.Contains(t, wrappedErr.Error(), "additional context")
}

func TestNewErr(t *testing.T) {
	err := NewErr("test error", http.StatusInternalServerError)
	var e *Error
	ok := errors.As(err, &e)
	require.True(t, ok)
	assert.Equal(t, "test error", e.Msg)
	assert.Equal(t, DefaultErrorCode, e.Code)
	assert.Equal(t, http.StatusInternalServerError, e.StatusCode)
}

func TestHandling(t *testing.T) {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		Handling(c, NewErr("test error"))
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "test error")
}

func TestReconstructStackTrace(t *testing.T) {
	err := NewErr("test error")
	trace, traced := ReconstructStackTrace(err, err)
	assert.NotEmpty(t, trace)
	assert.True(t, traced)
}
