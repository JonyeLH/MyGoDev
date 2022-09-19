package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc func(ctx *Context)) //添加method		//直接使用Routable来实现路由
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(write http.ResponseWriter, request *http.Request) {
		ctx := NewContext(write, request)
		handleFunc(ctx)
	})
}
func (s *sdkHttpServer) Start(address string) error {
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
