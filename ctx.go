package ctx

import (
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

func (c *Context) JSON(status int, data interface{}) error {
	return ngamux.JSONWithStatus(c.rw, status, data)
}

func (c *Context) String(status int, data string) error {
	return ngamux.StringWithStatus(c.rw, status, data)
}
