package main

import (
	"net/http"
)

type HandlerBasedMap struct {
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h *HandlerBasedMap) Key(method string, pattern string) string {
	return method + "#" + pattern
}
