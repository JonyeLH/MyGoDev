package main

import "fmt"

//全局变量的声明初始化，全局变量非必须使用
var a = "initial"
var b, c int = 1, 2
var d = true

func main() {
	fmt.Println(a)
	fmt.Println(b + c)
	fmt.Println(d)
	s := "GOGOGO" //局部变量可以使用这种方法声明并初始化，局部变量必须使用
	fmt.Println(s)
}
