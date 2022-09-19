package main

import "fmt"

func vals(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	A, B := vals(1, 2)
	fmt.Println("A is ", A)
	fmt.Println("B is ", B)

	_, C := vals(2, 3)
	fmt.Println("C is ", C)
}
