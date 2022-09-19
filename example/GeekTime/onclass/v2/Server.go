package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

// 结构体作为接口的方法接收器，最好都是用指针的形式
func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{ //当返回实际类型所实现的接口的时候，需要返回指针
		Name: name,
	}
}

//若是type A B形式，就不使用指针形式，但这种方法一般不使用
type Handle func()

func (h Handle) Hello() {

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}

	// 使用Context 处理json
	ctx := Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
	}

	resp := commonResponse{
		Data: 123,
	}

	err = ctx.OkJson(resp)
	if err != nil {
		fmt.Fprintf(w, "error %v ", err)
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
