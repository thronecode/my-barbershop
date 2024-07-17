package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseParams(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/?limit=10&offset=5&order=name&desc=true&total=true&filter1=value1&filter2=value2", nil)
		require.NoError(t, err)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = req

		params, err := ParseParams(c)
		require.NoError(t, err)

		assert.Equal(t, 10, params.Limit)
		assert.Equal(t, 5, params.Offset)
		assert.Equal(t, "name", params.OrderKey)
		assert.True(t, params.Desc)
		assert.True(t, params.Total)
		assert.Equal(t, []string{"value1"}, params.Filters["filter1"])
		assert.Equal(t, []string{"value2"}, params.Filters["filter2"])
	})

	t.Run("Default", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.NoError(t, err)

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = req

		params, err := ParseParams(c)
		require.NoError(t, err)

		assert.Equal(t, 15, params.Limit)
		assert.Equal(t, 0, params.Offset)
		assert.Equal(t, "", params.OrderKey) // default value
		assert.False(t, params.Desc)
		assert.False(t, params.Total)
		assert.Empty(t, params.Filters)
	})
}

func TestConvertFilters(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		type TestStruct struct {
			IntField    *int       `converter:"int_field"`
			StringField *string    `converter:"string_field"`
			TimeField   *time.Time `converter:"time_field"`
		}

		intValue := 42
		stringValue := "test"
		timeValue := time.Now()

		params := RequestParams{
			Filters: map[string][]string{},
		}

		testStruct := TestStruct{
			IntField:    &intValue,
			StringField: &stringValue,
			TimeField:   &timeValue,
		}

		err := params.ConvertFilters(&testStruct)
		require.NoError(t, err)

		assert.Equal(t, []string{"42"}, params.Filters["int_field"])
		assert.Equal(t, []string{"test"}, params.Filters["string_field"])
		assert.Equal(t, []string{timeValue.Format(time.RFC3339)}, params.Filters["time_field"])
	})

	t.Run("Empty", func(t *testing.T) {
		type EmptyStruct struct{}

		params := RequestParams{
			Filters: map[string][]string{},
		}

		emptyStruct := EmptyStruct{}

		err := params.ConvertFilters(&emptyStruct)
		assert.Error(t, err)
		assert.Equal(t, "no available fields", err.Error())
	})
}
