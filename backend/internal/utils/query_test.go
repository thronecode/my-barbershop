package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetColumns(t *testing.T) {
	type TestStruct struct {
		ID   int    `sql:"id" converter:"id"`
		Name string `sql:"name" converter:"name"`
	}

	// Testa sem o parâmetro `total`
	columns := GetColumns(TestStruct{}, nil)
	expected := []string{"id as id", "name as name"}
	assert.ElementsMatch(t, expected, columns)

	// Testa com o parâmetro `total` verdadeiro
	total := true
	columnsWithTotal := GetColumns(TestStruct{}, &total)
	expectedWithTotal := []string{"count(*)"}
	assert.ElementsMatch(t, expectedWithTotal, columnsWithTotal)
}

func TestScanStruct(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}

	v := reflect.ValueOf(&TestStruct{ID: 1, Name: "Test"})
	fields := scanStruct(v.Elem())

	assert.Len(t, fields, 2)

	assert.IsType(t, new(int), fields[0])
	assert.IsType(t, new(string), fields[1])
}
