package ctx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-must/must"
)

func TestHandler(t *testing.T) {
	t.Run("should returns same value using handler or not", func(t *testing.T) {
		must := must.New(t)

		notHandlerRW := httptest.NewRecorder()
		notHandlerR := httptest.NewRequest(http.MethodGet, "/", nil)
		err := func(w http.ResponseWriter, r *http.Request) error {
			c := New(w, r)
			return c.Res().Status(http.StatusOK).Text("aaa")
		}(notHandlerRW, notHandlerR)
		must.Nil(err)

		handlerRW := httptest.NewRecorder()
		handlerR := httptest.NewRequest(http.MethodGet, "/", nil)
		err = Handler(func(c *Context) error {
			return c.Res().Status(http.StatusOK).Text("aaa")
		})(handlerRW, handlerR)
		must.Nil(err)

		must.Equal(notHandlerRW.Result().StatusCode, handlerRW.Result().StatusCode)
		must.Equal(notHandlerRW.Result().Body, handlerRW.Result().Body)
	})
}
