package main

import "fmt"

//闭包，相当于就是匿名函数
//形式：func function_name() func() type{···}
//运行方法：
//1、在匿名函数最后添加()，调用时将自动运行
//func function_name() func() type{···}()
//2、在匿名函数赋值后，在赋值对象后跟个()
//Function := function_name()
//Function()

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())
}
