package ctx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-must/must"
	"github.com/ngamux/ngamux"
)

func TestNew(t *testing.T) {
	t.Run("should returns ctx.Context instance", func(t *testing.T) {
		must := must.New(t)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		expected := &Context{
			ngamux.Req(r),
			ngamux.Res(w),
		}

		actual := New(w, r)

		must.NotNil(actual)
		must.Equal(expected, actual)
		must.Equal(expected.req, actual.req)
		must.Equal(expected.res, actual.res)
	})
}

func BenchmarkNew(b *testing.B) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			New(w, r)
		}
	}
}
