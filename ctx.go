package ctx

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

type Context struct {
	req *ngamux.Request
	res *ngamux.Response
}

func New(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ngamux.Req(r),
		ngamux.Res(w),
	}
}

// REQUEST

func (c *Context) Req() *ngamux.Request {
	return c.req
}

// RESPONSE

func (c *Context) Res() *ngamux.Response {
	return c.res
}
