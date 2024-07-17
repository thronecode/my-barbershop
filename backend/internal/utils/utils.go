package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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

			dstValue := elemDst.Field(d)

			if srcField.Type != dstField.Type {
				return fmt.Errorf("type mismatch: %s != %s", srcField.Type, dstField.Type)
			} else {
				dstValue.Set(elemSrc.Field(s))
			}
		}
	}

	return nil
}
