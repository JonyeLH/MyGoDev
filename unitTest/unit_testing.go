package main

import "fmt"

/*
- 文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
- 必须import testing这个包
- 所有的测试用例函数必须是Test开头
- 测试用例会按照源代码中写的顺序依次执行
- 测试函数TestXxx()的参数是testing.T，可以使用该类型来记录错误或者是测试状态
- 测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
- 函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。
参考链接：https://www.cnblogs.com/sunylat/p/6387356.html
*/

type TestStruct struct {
	Id   string
	Name string
}

//这是一个结构函数，返回两个值，输入为方法接收者
func (ts *TestStruct) StructFunc(idIn, nameIn string) (idOut, nameOut string) {

	ts.Id = idIn
	ts.Name = nameIn

	return ts.Id, ts.Name
}

//这是另外一个结构函数，返回结构，输入为方法接收者
func (ts *TestStruct) StructFunc2(idIn, nameIn string) TestStruct {
	ts.Id = idIn
	ts.Name = nameIn
	return *ts
}

//这是普通函数，函数名前面没有方法接收者
func OrdinaryFunc(input1, input2 string) (outPut string) {
	temp := input1 + input2
	return temp
}

func Add(a, b int) int {
	fmt.Println(a + b)
	return a + b
}
