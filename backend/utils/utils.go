package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

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
