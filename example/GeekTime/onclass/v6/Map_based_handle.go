package main

import "net/http"

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}
type Handler interface {
	http.Handler
	Routable
}

type HandlerBasedMap struct {
	handlers map[string]func(ctx *Context)
}

func (s *HandlerBasedMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	key := s.Key(method, pattern)
	s.handlers[key] = handleFunc
}

func (h *HandlerBasedMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.Key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedMap) Key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedMap{} //用于判断是否真的 创建HandlerBasedMap
func NewHandlerBasedMap() Handler {
	return &HandlerBasedMap{
		handlers: make(map[string]func(ctx *Context), 128), //这里可以预估好创建的路由数目
	}
}
