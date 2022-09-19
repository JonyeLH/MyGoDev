package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) //这里是值传递，在别的地方修改不会影响原始数据
	fmt.Println("zeroval:", i)

	zeroptr(&i) //若是传递指针，在别的地方修改就是对变量的内存空间修改，就会修改原始数据
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
