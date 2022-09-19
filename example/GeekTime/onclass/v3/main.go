package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

func main() {
	server := NewHttpServer("Server_test")
	//server.Route("/", home)
	//server.Route("/user", user)
	//server.Route("/user/create", createUser)
	//server.Route("/order", order)
	//server.Route("/user/signUp", SignUp)
	server.Route("/user/signUp", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err) //如果进入错误，panic直接停止main函数，panic是一种不可恢复的错误判断，一旦使用了，后面的代码就不会执行包括goroute
		//在一旦错误失败的出现，整个代码就没意义执行的场景，使用panic
		//一般使用error即可
	}
}
