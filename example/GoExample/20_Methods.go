package main

import "fmt"

//方法接收器，建议使用指针
type rect struct {
	width, length int
}

func (r *rect) area() int {
	return r.length * r.width
}

func (r rect) perim() int {
	return 2*r.length + 2*r.width
}

func main() {
	r := rect{
		width:  10,
		length: 5,
	}
	fmt.Println("area= ", r.area())
	fmt.Println("perim= ", r.perim())

	rp := &r
	fmt.Println("area= ", rp.area())
	fmt.Println("perim= ", rp.perim())
}
