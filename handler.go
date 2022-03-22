package ctx

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

type CtxHandler func(c *Context) error

func Handler(handler CtxHandler) ngamux.Handler {
	return func(rw http.ResponseWriter, r *http.Request) error {
		c := New(rw, r)
		return handler(c)
	}
}
