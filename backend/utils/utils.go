package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// MaxLimit defines query max value to field limit
const MaxLimit int = 100000

// RequestParams used to store the parameters of a request
type RequestParams struct {
	Filters  map[string][]string
	OrderKey string
	Limit    uint64
	Offset   uint64
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
	params.Limit = uint64(min(limit, MaxLimit))

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		return params, err
	}
	params.Offset = uint64(offset)

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
		var (
			field  = structureRt.Field(i)
			values []string
		)

		key := structureRt.Type().Field(i).Tag.Get(converterTag)
		if field.IsNil() || key == "" {
			continue
		}

		switch value := field.Interface().(type) {
		case *int, *int8, *int32, *int64, *bool, *float32, *float64, *string:
			values = append(values, fmt.Sprint(field.Elem().Interface()))

		case []int, []int8, []int32, []int64, []bool, []float32, []float64, []string:
			rt := reflect.ValueOf(value)

			for j := 0; j < rt.Len(); j++ {
				switch valor := rt.Index(j).Interface().(type) {
				case *time.Time:
					values = append(values, valor.Format(time.RFC3339))
				default:
					values = append(values, fmt.Sprint(rt.Index(j).Interface()))
				}
			}

		case *time.Time:
			values = append(values, value.Format(time.RFC3339))

		default:
			continue
		}

		p.Filters[key] = values
	}

	return nil
}

// IsStringInSlice checks if a string is in a slice
func IsStringInSlice(a string, list ...string) bool {
	for _, b := range list {
		if strings.EqualFold(b, a) {
			return true
		}
	}

	return false
}

// ConvertStruct converts struct data in to a destination data out
func ConvertStruct(in interface{}, out interface{}) error {
	tagName := "converter"
	elemSrc := reflect.ValueOf(in).Elem()
	elemDst := reflect.ValueOf(out).Elem()

	if elemSrc.Kind() != reflect.Struct || elemDst.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	if elemSrc.NumField() == 0 || elemDst.NumField() == 0 {
		return errors.New("no available fields")
	}

	for s := 0; s < elemSrc.NumField(); s++ {
		srcField := elemSrc.Type().Field(s)

		srcKey := srcField.Tag.Get(tagName)
		if srcKey == "" {
			continue
		}
		if elemSrc.Field(s).Kind() == reflect.Ptr && elemSrc.Field(s).IsNil() {
			continue
		}

		for d := 0; d < elemDst.NumField(); d++ {
			dstField := elemDst.Type().Field(d)

			dstKey := dstField.Tag.Get(tagName)
			if dstKey == "" || dstKey != srcKey {
				continue
			}

			if !elemDst.Field(d).CanSet() {
				return errors.New("destination is not settable")
			}

			if srcField.Type != dstField.Type {
				var (
					tSrc, tDst string
				)

				if srcField.Type.String()[0] == '*' {
					tSrc = srcField.Type.Elem().String()
				} else {
					tSrc = srcField.Type.String()
				}

				if dstField.Type.String()[0] == '*' {
					tDst = dstField.Type.Elem().String()
				} else {
					tDst = dstField.Type.String()
				}

				if dstField.Type.String() == "*time.Time" {
					tDst = dstField.Type.String()
				}

				val, err := ConvertStringToTime(elemSrc.Field(s).Interface(), tSrc, tDst)
				if err != nil {
					return fmt.Errorf("was not possible to convert field %s: %s", srcField.Name, err.Error())
				}

				v := reflect.ValueOf(val)
				elemDst.Field(d).Set(v)
			} else {
				elemDst.Field(d).Set(elemSrc.Field(s))
			}
		}
	}

	return nil
}

// ConvertStringToTime converts strings in time pointer
func ConvertStringToTime(value interface{}, from, to string) (output interface{}, err error) {
	if from == "string" {
		if to == "*time.Time" {
			var val string
			if v, ok := value.(*string); ok {
				if v == nil {
					return nil, nil
				}
				val = *v
			} else if v, ok := value.(string); ok {
				val = v
			} else {
				return nil, nil
			}
			return time.Parse(time.RFC3339, val)
		}
	}

	return nil, errors.New("feature not supported")
}
