package ctx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-must/must"
	"github.com/ngamux/ngamux"
)

func TestMux(t *testing.T) {
	t.Run("should return CtxMux", func(t *testing.T) {
		must := must.New(t)

		mux := ngamux.New()
		expected := &CtxMux{mux}
		actual := Mux(mux)

		must.NotNil(actual)
		must.Equal(expected, actual)
	})
}

func TestCtxMuxGet(t *testing.T) {
	t.Run("should append route with Get method", func(t *testing.T) {
		must := must.New(t)

		cmux := Mux(ngamux.New())

		expectedStatus := 200
		expectedBody := "hore"
		cmux.Get("/", func(c *Context) error { return c.String(expectedStatus, expectedBody) })

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		cmux.mux.ServeHTTP(rw, r)

		actualStatus := rw.Result().StatusCode
		actualBody := strings.ReplaceAll(rw.Body.String(), "\n", "")

		must.Equal(expectedStatus, actualStatus)
		must.Equal(expectedBody, actualBody)
	})
}

func TestCtxMuxPost(t *testing.T) {
	t.Run("should append route with Post method", func(t *testing.T) {
		must := must.New(t)

		cmux := Mux(ngamux.New())

		expectedStatus := 200
		expectedBody := "hore"
		cmux.Post("/", func(c *Context) error { return c.String(expectedStatus, expectedBody) })

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		cmux.mux.ServeHTTP(rw, r)

		actualStatus := rw.Result().StatusCode
		actualBody := strings.ReplaceAll(rw.Body.String(), "\n", "")

		must.Equal(expectedStatus, actualStatus)
		must.Equal(expectedBody, actualBody)
	})
}

func TestCtxMuxPut(t *testing.T) {
	t.Run("should append route with Put method", func(t *testing.T) {
		must := must.New(t)

		cmux := Mux(ngamux.New())

		expectedStatus := 200
		expectedBody := "hore"
		cmux.Put("/", func(c *Context) error { return c.String(expectedStatus, expectedBody) })

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/", nil)
		cmux.mux.ServeHTTP(rw, r)

		actualStatus := rw.Result().StatusCode
		actualBody := strings.ReplaceAll(rw.Body.String(), "\n", "")

		must.Equal(expectedStatus, actualStatus)
		must.Equal(expectedBody, actualBody)
	})
}

func TestCtxMuxPatch(t *testing.T) {
	t.Run("should append route with Patch method", func(t *testing.T) {
		must := must.New(t)

		cmux := Mux(ngamux.New())

		expectedStatus := 200
		expectedBody := "hore"
		cmux.Patch("/", func(c *Context) error { return c.String(expectedStatus, expectedBody) })

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/", nil)
		cmux.mux.ServeHTTP(rw, r)

		actualStatus := rw.Result().StatusCode
		actualBody := strings.ReplaceAll(rw.Body.String(), "\n", "")

		must.Equal(expectedStatus, actualStatus)
		must.Equal(expectedBody, actualBody)
	})
}

func TestCtxMuxDelete(t *testing.T) {
	t.Run("should append route with Delete method", func(t *testing.T) {
		must := must.New(t)

		cmux := Mux(ngamux.New())

		expectedStatus := 200
		expectedBody := "hore"
		cmux.Delete("/", func(c *Context) error { return c.String(expectedStatus, expectedBody) })

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		cmux.mux.ServeHTTP(rw, r)

		actualStatus := rw.Result().StatusCode
		actualBody := strings.ReplaceAll(rw.Body.String(), "\n", "")

		must.Equal(expectedStatus, actualStatus)
		must.Equal(expectedBody, actualBody)
	})
}

func TestCtxMuxAll(t *testing.T) {
	t.Run("should append route with All method", func(t *testing.T) {
		must := must.New(t)

		cmux := Mux(ngamux.New())

		expectedStatus := 200
		expectedBody := "hore"
		cmux.All("/", func(c *Context) error { return c.String(expectedStatus, expectedBody) })

		rw := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		cmux.mux.ServeHTTP(rw, r)

		actualStatus := rw.Result().StatusCode
		actualBody := strings.ReplaceAll(rw.Body.String(), "\n", "")

		must.Equal(expectedStatus, actualStatus)
		must.Equal(expectedBody, actualBody)
	})
}
