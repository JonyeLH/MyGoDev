package main

import "net/http"

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}
type Handler interface {
	ServerHTTP(c *Context)
	Routable
}

type HandlerBasedMap struct {
	handlers map[string]func(ctx *Context)
}

func (s *HandlerBasedMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	key := s.Key(method, pattern)
	s.handlers[key] = handleFunc
}

func (h *HandlerBasedMap) ServerHTTP(c *Context) {
	key := h.Key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
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
