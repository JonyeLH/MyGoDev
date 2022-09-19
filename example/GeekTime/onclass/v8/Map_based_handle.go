package main

import (
	"net/http"
)
//type handlerFunc func(ctx *Context)
type Routable interface {
	Route(method string, pattern string, handleFunc handlerFunc)
}
type Handler interface {
	ServerHTTP(c *Context)
	Routable
}

type HandlerBasedMap struct {
	handlers map[string]handlerFunc
	//handlers sync.Map
}

func (s *HandlerBasedMap) Route(method string, pattern string, handleFunc handlerFunc) {
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
		handlers: make(map[string]handlerFunc, 128), //这里可以预估好创建的路由数目
	}
}
