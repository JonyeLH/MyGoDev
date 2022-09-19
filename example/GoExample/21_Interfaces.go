package main

//接口，结合结构体使用方法接收器：goland操作，在结构体中使用Ctrl+I，选择相应接口，再次选中接收器的类型

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64  //面积
	perim() float64 //周长
}
type rec struct {
	width, height float64 //长宽
}

func (r *rec) area() float64 {
	return r.height * r.width
}

func (r *rec) perim() float64 {
	return 2*r.height + 2*r.width
}

type circle struct {
	radius float64 //定义半径
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rec{width: 3, height: 4}
	c := circle{radius: 5}

	measure(&r)
	measure(c)
}
