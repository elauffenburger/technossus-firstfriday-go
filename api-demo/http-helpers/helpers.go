package helpers

import "net/http"

type HandlerAdapter struct {
    fn func(http.ResponseWriter, *http.Request)
}

func (h HandlerAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    h.fn(w, req)
}

func MakeHandler(fn func(w http.ResponseWriter, req *http.Request)) *HandlerAdapter {
    return &HandlerAdapter{fn: fn}
}

type HandlerFn func(w http.ResponseWriter, req *http.Request)

func (t HandlerFn) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    t(w, req)
}

func (t HandlerFn) ToHandler() http.Handler {
    return http.Handler(t)
}
