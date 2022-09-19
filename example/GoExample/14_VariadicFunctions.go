package main

import "fmt"

//... 表示可以读入任意个输入参数
func sum(nums ...int) {
	fmt.Println("nums= ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total= ", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	num := []int{1, 2, 3, 4}
	sum(num...) //表示输入任意个输入参数
}
