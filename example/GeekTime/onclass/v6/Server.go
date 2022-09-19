package main

import "net/http"

//采用filter实现路由串接
type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handleFunc)
}
func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedMap()
	var root Filter = func(ctx *Context) {
		handler.ServeHTTP(ctx.W, ctx.R) //这里的ServeHTTP函数跳转有问题，不是直接调至Map_based_handle的函数，而是原生的ServeHTTP
	}
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{ //当返回实际类型所实现的接口的时候，需要返回指针
		Name:    name,
		handler: handler,
		root:    root,
	}
}

func SignUp(ctx *Context) {
	resp := &signUpReq{}
	err := ctx.ReadJson(resp)
	if err != nil {
		err := ctx.BadRequestJson(err)
		if err != nil {
			return
		}
	}

	rep := commonResponse{
		Data: 123,
	}
	err = ctx.OkJson(rep)
	if err != nil {
		err := ctx.BadRequestJson(err)
		if err != nil {
			return
		}
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
