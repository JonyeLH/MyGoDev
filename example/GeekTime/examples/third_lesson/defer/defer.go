package main

import "fmt"

func main() {
	//defer定义的函数是在return之前执行的函数，甚至可在panic后执行，是一种栈的实现，下例中是由下往上执行
	defer func() {
		fmt.Println("aaa")
	}()

	defer func() {
		fmt.Println("bbb")
	}()

	defer func() {
		fmt.Println("ccc")
	}()
}
