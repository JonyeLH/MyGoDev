package main

type AAA struct {
	Name string
	Age int
}

func (c *AAA) ChangeName(name string) {
	c.Name = name
}

func (c *AAA) ChangeAge(age int) {
	c.Age = age
}

func (AAA) ChangeNoting(){

}
