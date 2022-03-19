package ctx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-must/must"
)

func TestNew(t *testing.T) {
	t.Run("should returns ctx.Context instance", func(t *testing.T) {
		must := must.New(t)

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		expected := &Context{
			rw,
			r,
		}

		actual := New(rw, r)

		must.NotNil(actual)
		must.Equal(expected, actual)

	})
}

func TestGetJSON(t *testing.T) {
	expected := map[string]int{"id": 1}
	dataJSON, _ := json.Marshal(expected)

	rw := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(dataJSON)))
	c := New(rw, r)

	t.Run("should returns JSON from body", func(t *testing.T) {
		must := must.New(t)
		var actual map[string]int
		err := c.GetJSON(&actual)

		must.Nil(err)
		must.NotNil(actual)
		must.Equal(expected, actual)
	})
}
