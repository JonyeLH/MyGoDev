package main

import "net/http"

///////////////////// 添加Method /////////////////////
type Server interface {
	Route(method string, pattern string, handleFunc func(ctx *Context)) //添加method		//直接使用Routable来实现路由
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedMap //添加HandlerBasedMap		//强依赖handler、HandlerBasedMap
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	key := s.handler.Key(method, pattern) //强依赖handler、HandlerBasedMap
	s.handler.handlers[key] = handleFunc  //强依赖handler、HandlerBasedMap
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{ //当返回实际类型所实现的接口的时候，需要返回指针
		Name: name,
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
