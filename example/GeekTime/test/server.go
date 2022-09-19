///////////////////// V2 ////////////////////////////
//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type Server interface {
//	Route(pattern string, handlerFunc http.HandlerFunc)
//	Start(address string) error
//}
//
//type sdkHttpServer struct{
//	Name string
//}
//
//func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
//	http.HandleFunc(pattern, handlerFunc)
//}
//
//func (s *sdkHttpServer) Start(address string) error{
//	return http.ListenAndServe(address,nil)
//}
//
//func NewHttpServer(name string) Server{
//	return &sdkHttpServer{
//		Name: name,
//	}
//}
//
//func SignUp(write http.ResponseWriter, request *http.Request) {
//	resp := &signUpReq{}
//	ctx := Context{
//		W: write,
//		R: request,
//	}
//	err := ctx.ReadJson(resp)
//	if err != nil{
//		fmt.Fprintf(write, "error %v", err)
//	}
//
//	rep := &commonResponse{
//		Data: 1234,
//	}
//	err = ctx.OkJson(rep)
//	if err != nil {
//		fmt.Fprintf(write, "error %v",err)
//	}
//}
//
//type signUpReq struct {
//	Email 				string	`json:"email"`
//	Password 			string	`json:"password"`
//	ConfirmedPassword	string	`json:"confirmed_password"`
//}
//
//type commonResponse struct {
//	BizCode 		int	`json:"biz_code"`
//	Msg				string	`json:"msg"`
//	Data 			interface{}	`json:"data"`
//}
///////////////////// V2 ////////////////////////////

///////////////////// V3 ////////////////////////////
//package main
//
//import (
//	"net/http"
//)
//
//type Server interface {
//	Route(pattern string, handlerFunc func(ctx *Context))
//	Start(address string) error
//}
//
//type sdkHttpServer struct{
//	Name string
//}
//
//func (s *sdkHttpServer) Route(pattern string,
//	handlerFunc func(ctx *Context)) {
//	http.HandleFunc(pattern,
//		func(writer http.ResponseWriter,
//			request *http.Request) {
//			ctx := NewContext(writer, request)
//			handlerFunc(ctx)
//		})
//}
//
//func (s *sdkHttpServer) Start(address string) error{
//	return http.ListenAndServe(address,nil)
//}
//
//func NewHttpServer(name string) Server{
//	return &sdkHttpServer{
//		Name: name,
//	}
//}
//
//func SignUp(ctx *Context) {
//	resp := &signUpReq{}
//
//	err := ctx.ReadJson(resp)
//	if err != nil{
//		err = ctx.BadRequestJson(err)
//	}
//
//	rep := &commonResponse{
//		Data: 1234,
//	}
//	err = ctx.OkJson(rep)
//	if err != nil {
//		err = ctx.BadRequestJson(err)
//	}
//}
//
//type signUpReq struct {
//	Email 				string	`json:"email"`
//	Password 			string	`json:"password"`
//	ConfirmedPassword	string	`json:"confirmed_password"`
//}
//
//type commonResponse struct {
//	BizCode 		int	`json:"biz_code"`
//	Msg				string	`json:"msg"`
//	Data 			interface{}	`json:"data"`
//}
/////////////////// V3 ////////////////////////////

/////////////////// V4 ////////////////////////////
package main

import (
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedMap
}

func (s *sdkHttpServer) Route(
	method string,
	pattern string,
	handlerFunc func(ctx *Context)) {
	key := s.handler.Key(method, pattern)
	s.handler.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

func SignUp(ctx *Context) {
	resp := &signUpReq{}

	err := ctx.ReadJson(resp)
	if err != nil {
		err = ctx.BadRequestJson(err)
	}

	rep := &commonResponse{
		Data: 1234,
	}
	err = ctx.OkJson(rep)
	if err != nil {
		err = ctx.BadRequestJson(err)
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
