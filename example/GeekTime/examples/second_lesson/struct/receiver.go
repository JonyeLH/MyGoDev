package main

import "fmt"

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")
	up.ChangeAge(120)

	fmt.Printf("%v \n", up)
}

type User struct {
	Name string
	Age  int
}

// 结构体接收器
// 当我们提供给外的接口不想让别人修改我们结构体数据时，使用结构体接收器，
// 其实就是值复制，无论其他地方怎么改，都不会影响结构体的数据
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器	遇事不决用指针
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

//type A B类型则使用结构体
type handle func()

func (h handle) Hello() {

}
