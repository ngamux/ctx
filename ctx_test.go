package ctx

import (
	"net/http"
	"net/http/httptest"
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
