package ctx

import (
	"context"
	"net/http"

	"github.com/ngamux/ngamux"
)

type Context struct {
	rw http.ResponseWriter
	r  *http.Request
}

func New(rw http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		rw,
		r,
	}
}

// REQUEST

func (c *Context) Request() *http.Request {
	return c.r
}

func (c *Context) Context() context.Context {
	return c.r.Context()
}

func (c *Context) GetJSON(store interface{}) error {
	return ngamux.GetJSON(c.r, store)
}

func (c *Context) GetParam(key string) string {
	return ngamux.GetParam(c.r, key)
}

func (c *Context) GetQuery(key string) string {
	return ngamux.GetQuery(c.r, key)
}

// RESPONSE

func (c *Context) Response() http.ResponseWriter {
	return c.rw
}

func (c *Context) JSON(status int, data interface{}) error {
	return ngamux.JSONWithStatus(c.rw, status, data)
}

func (c *Context) String(status int, data string) error {
	return ngamux.StringWithStatus(c.rw, status, data)
}
