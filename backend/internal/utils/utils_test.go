package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsStringInSlice(t *testing.T) {
	assert.True(t, IsStringInSlice("test", "test", "other"))
	assert.False(t, IsStringInSlice("test", "other", "another"))
	assert.True(t, IsStringInSlice("TEST", "test", "other")) // Test case insensitive
}

func TestConvertStruct(t *testing.T) {
	type Source struct {
		ID   int        `converter:"id"`
		Date *time.Time `converter:"date"`
		Name string     `converter:"name"`
	}

	type Destination struct {
		ID   int        `converter:"id"`
		Date *time.Time `converter:"date"`
		Name string     `converter:"name"`
	}

	date := time.Now()
	source := &Source{
		ID:   1,
		Date: &date,
		Name: "Test Name",
	}

	var destination Destination
	err := ConvertStruct(source, &destination)
	require.NoError(t, err)
	assert.Equal(t, source.ID, destination.ID)
	assert.Equal(t, source.Name, destination.Name)
	assert.NotNil(t, destination.Date)
	assert.Equal(t, *source.Date, *destination.Date)
}

func TestConvertStructMismatch(t *testing.T) {
	type Source struct {
		ID   int        `converter:"id"`
		Date *time.Time `converter:"date"`
		Name string     `converter:"name"`
	}

	type Destination struct {
		ID   int     `converter:"id"`
		Date *string `converter:"date"`
		Name string  `converter:"name"`
	}

	date := time.Now()
	source := &Source{
		ID:   1,
		Date: &date,
		Name: "Test Name",
	}

	var destination Destination
	err := ConvertStruct(source, &destination)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "type mismatch")
}
