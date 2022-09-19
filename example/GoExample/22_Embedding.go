package main

import "fmt"

//组合：结构体包另外个结构体，接口也可以
type base struct {
	num int
}

func (b *base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	b := base{
		num: 13,
	}
	fmt.Println("result describe = ", b.describe())

	co := container{
		base: base{
			num: 123,
		},
		str: "Hello",
	}
	fmt.Printf("co = {num: %v, str = %v}\n", co.num, co.str)
	fmt.Println("also num ", co.base.num)
	fmt.Println("also func ", co.describe())

	type describer interface {
		describe() string
	}
	var coo describer = &co
	fmt.Println("describe = ", coo.describe())
}
