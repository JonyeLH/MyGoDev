package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20}) //{Bob 20}

	fmt.Println(person{name: "Alice", age: 30}) //{Alice 30}

	fmt.Println(person{name: "Fred"}) //{Fred 0}

	fmt.Println(&person{name: "Ann", age: 40}) //&{Ann 40}

	fmt.Println(newPerson("Jon")) //&{Jon 42}

	//fmt.Println(newPerson("Jon", 26))			//&{Jon 42}

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) //Sean

	sp := &s
	fmt.Println(sp.age) //50

	sp.age = 51         //sp是一个指针
	fmt.Println(sp.age) //51
}
