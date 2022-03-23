package ctx

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

type CtxMux struct {
	mux *ngamux.Ngamux
}

func Mux(mux *ngamux.Ngamux) *CtxMux {
	return &CtxMux{
		mux,
	}
}

func (m *CtxMux) Get(path string, handler CtxHandler) {
	m.mux.Get(path, Handler(handler))
}

func (m *CtxMux) Post(path string, handler CtxHandler) {
	m.mux.Post(path, Handler(handler))
}

func (m *CtxMux) Put(path string, handler CtxHandler) {
	m.mux.Put(path, Handler(handler))
}

func (m *CtxMux) Patch(path string, handler CtxHandler) {
	m.mux.Patch(path, Handler(handler))
}

func (m *CtxMux) Delete(path string, handler CtxHandler) {
	m.mux.Delete(path, Handler(handler))
}

func (m *CtxMux) All(path string, handler CtxHandler) {
	m.mux.All(path, Handler(handler))
}

func (m *CtxMux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(rw, r)
}
