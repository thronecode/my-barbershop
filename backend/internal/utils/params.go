package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// MaxLimit defines query max value to field limit
const MaxLimit int = 100000

// RequestParams used to store the parameters of a request
type RequestParams struct {
	Filters  map[string][]string
	OrderKey string
	Limit    int
	Offset   int
	Desc     bool
	Total    bool
}

// ParseParams receives the gin.Context and parse the query params for the request
func ParseParams(c *gin.Context) (RequestParams, error) {
	params := RequestParams{}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "15"))
	if err != nil {
		return params, err
	}

	if limit <= 0 {
		limit = MaxLimit
	}
	params.Limit = min(limit, MaxLimit)

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		return params, err
	}
	params.Offset = offset

	params.OrderKey = c.DefaultQuery("order", "")

	if params.Desc, err = strconv.ParseBool(c.DefaultQuery("desc", "false")); err != nil {
		return params, err
	}

	if params.Total, err = strconv.ParseBool(c.DefaultQuery("total", "false")); err != nil {
		return params, err
	}

	params.Filters = map[string][]string{}
	for k, v := range c.Request.URL.Query() {
		if IsStringInSlice(k, "limit", "offset", "order", "desc", "total") {
			continue
		}

		if len(v) > 0 {
			params.Filters[k] = append(params.Filters[k], v...)
		}
	}

	return params, nil
}

// ConvertFilters populates filters using fields from the
// received structure that have the "converter" tag
func (p *RequestParams) ConvertFilters(structure any) error {
	const converterTag = "converter"

	if p.Filters == nil {
		p.Filters = map[string][]string{}
	}

	structureRt := reflect.ValueOf(structure).Elem()

	if structureRt.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	if structureRt.NumField() == 0 {
		return errors.New("no available fields")
	}

	for i := 0; i < structureRt.NumField(); i++ {
		field := structureRt.Field(i)
		tag := structureRt.Type().Field(i).Tag.Get(converterTag)

		if tag == "" {
			continue
		}

		var values []string

		// Handle pointer types
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}
			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Bool, reflect.Float32, reflect.Float64, reflect.String:
			values = append(values, fmt.Sprint(field.Interface()))
		case reflect.Slice:
			for j := 0; j < field.Len(); j++ {
				val := field.Index(j).Interface()
				switch val := val.(type) {
				case *time.Time:
					values = append(values, val.Format(time.RFC3339))
				default:
					values = append(values, fmt.Sprint(val))
				}
			}
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				values = append(values, field.Interface().(time.Time).Format(time.RFC3339))
			}
		default:
			continue
		}

		p.Filters[tag] = values
	}

	return nil
}
