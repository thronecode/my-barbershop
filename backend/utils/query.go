package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Masterminds/squirrel"
)

// GetColumns returns the columns to be selected in the query
func GetColumns(structure any, total *bool) []string {
	if total != nil && *total {
		return []string{"count(*)"}
	}

	var columns []string
	val := reflect.ValueOf(structure)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		sqlTag := field.Tag.Get("sql")
		converterTag := field.Tag.Get("converter")

		if sqlTag != "" && converterTag != "" {
			columns = append(columns, fmt.Sprintf("%s as %s", sqlTag, converterTag))
		}
	}

	return columns
}

// MakePaginatedList makes a paginated list
func MakePaginatedList(structure any, query *squirrel.SelectBuilder, params *RequestParams) (any, *bool, *int, error) {
	if query == nil {
		return nil, nil, nil, errors.New("query is nil")
	}

	if params.Total {
		var total int
		err := query.QueryRow().Scan(&total)
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, nil, &total, nil
	}

	*query = query.
		Limit(uint64(params.Limit + 1)).
		Offset(uint64(params.Offset))

	if params.OrderKey != "" {
		if params.OrderKey != "" {
			order := "asc"
			if params.Desc {
				order = "desc"
			}
			*query = query.OrderBy(params.OrderKey + " " + order)
		}
	}

	rows, err := query.Query()
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()

	var (
		data  = reflect.New(reflect.SliceOf(reflect.TypeOf(structure))).Elem()
		count = 0
	)

	for rows.Next() {
		count++
		elem := reflect.New(reflect.TypeOf(structure)).Elem()
		err = rows.Scan(scanStruct(elem)...)
		if err != nil {
			return nil, nil, nil, err
		}

		data = reflect.Append(data, elem)
	}

	var next bool
	if count > int(params.Limit) {
		next = true
		data = data.Slice(0, params.Limit)
	}

	return data.Interface(), &next, nil, nil
}

func scanStruct(v reflect.Value) []any {
	var fields []interface{}
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanSet() {
			fields = append(fields, v.Field(i).Addr().Interface())
		}
	}
	return fields
}
