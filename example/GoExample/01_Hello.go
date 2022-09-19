package main

import "fmt"

func main() {
	//fmt.Println(a ...interface{})					//控制台打印任意类型的输出。不管输出类型，始终在输出之间添加空格,并换行
	//fmt.Printf(format string, a ...interface{})	//根据format参数，控制台打印字符串类型的输出。参数之间不添加空格，不换行
	//fmt.Print(a ...interface{})					//根据默认的格式化参数，控制台打印任意类型的输出。输出类型都不是字符串时，各个输出之间添加空格，不会换行

	//fmt.Sprintln(a ...interface{})					//采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符，会换行
	//fmt.Sprintf(format string, a ...interface{})		//根据format格式化返回生成的string。参数之间不添加空格，不换行
	//fmt.Sprint(a ...interface{})						//采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格，不换行

	//
	//fmt.Fprintln(w io.Writer, a ...interface{})					//采用默认格式将其参数格式化并写入 w 。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符，会换行
	//fmt.Fprintf(w io.Writer, format string, a ...interface{})		//根据format格式化生成的string写入w。参数之间不添加空格，不换行
	//fmt.Fprint(w io.Writer, a ...interface{})						//采用默认格式将其参数格式化并写入 w 。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格，不换行

	a := 10
	s := "hello world"
	p := Person{Name: "wohu", Age: 25}
	c := []int{1, 2, 3, 4}

	fmt.Printf("p %%v is %v\n", p)   // p %v is {wohu 25}	默认值打印
	fmt.Printf("p %%+v is %+v\n", p) // p %+v is {Name:wohu Age:25}	值打印，并输出结构体时会添加字段名
	fmt.Printf("p %%#v is %#v\n", p) // p %#v is main.Person{Name:"wohu", Age:25}
	fmt.Printf("p type is %T\n", p)  // p type is main.Person

	fmt.Printf("a %%#v is %#v, a type is %T\n", a, a) // a %#v is 10, a type is int
	fmt.Printf("s %%#v is %#v, s type is %T\n", s, s) // s %#v is "hello world", s type is string
	fmt.Printf("c %%v is %v, c type is %T\n", c, c)   // c %v is [1 2 3 4], c type is []int
	fmt.Printf("c %%#v is %#v, c type is %T\n", c, c) // c %#v is []int{1, 2, 3, 4}, c type is []int

}

type Person struct {
	Name string
	Age  int
}
