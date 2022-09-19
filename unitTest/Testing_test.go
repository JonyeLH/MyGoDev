package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"net/url"
	"reflect"
	"testing"
)


func TestStructs(t *testing.T) {
	up := &AAA{
		Name: "DDD",
		Age: 10,
	}
	up.ChangeName("xiaoming")
	fmt.Println(up.Name)
	up.ChangeAge(100)
	fmt.Println(up.Age)
}


func TestCall(t *testing.T) {
	_ = call("http://localhost:8080/employee", "POST")
}


//获取当前目录下的文件
func TestReadFile(t *testing.T) {
	ReadFile()
}


//url中文转码和反转码
func Test_Escape(t *testing.T) {
	//url编码
	str := "中文-_."
	unstr := "%2f"
	fmt.Printf("url.QueryEscape:%s", url.QueryEscape(str))
	fmt.Println()
	s, _ := url.QueryUnescape(unstr)
	fmt.Printf("url.QueryUnescape:%s", s)
	fmt.Println()
	//字符转码
	hstr := "<"
	hunstr := "&lt"
	fmt.Printf("html.EscapeString:%s", html.EscapeString(hstr))
	fmt.Println()
	fmt.Printf("html.UnescapeString:%s", html.UnescapeString(hunstr))
	fmt.Println()
}

func TestXmlRead(t *testing.T) {
	XmlLoad()
}

func TestXml(t *testing.T) {
	XmlCreate()
}

func TestXlsx(t *testing.T) {
	CreateXlsx()
}

func TestTime(t *testing.T) {
	TimeStamp()
}

func TestType(t *testing.T) {
	var x int32 = 20
	fmt.Println("type:", reflect.TypeOf(x))
}

const (
	originFilePath = "D:\\go_test\\go_test\\go_base\\my_tet\\"
	targetFilePath = "D:\\go_test\\go_test\\go_base\\my_tet\\"
)

func CompressTest(t *testing.T) {
	if err := CompressZip(originFilePath, targetFilePath); err != nil {
		fmt.Println("压缩异常")
	}
}

func TestGin(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//StructFunc的测试方法
func TestStructFunc(t *testing.T) {

	//声明结构函数输入的变量
	var (
		idIn   string
		nameIn string
	)

	//给结构函数输入的变量赋值
	idIn = "inputId"
	nameIn = "inputName"

	//得到结构
	ts := &TestStruct{}

	//调用结构函数1
	idOut, nameOut := ts.StructFunc(idIn, nameIn)

	if idOut == idIn && nameOut == nameIn {
		t.Log("测试通过！")
	} else {
		t.Error("函数执行错误")
	}

}

/*
goconvey使用基本的方法
*/
//func TestAdd(t *testing.T) {
//	Convey("测试add方法", t, func() {
//		So(Add(2, 3), ShouldEqual, 5)
//	})
//}

//Json Decoder 使用方法
func TestDecoder(t *testing.T) {
	Decoder()
	fmt.Println("完成调用")
}

//find file path
func TestFind(t *testing.T) {
	//FindFile()
	FindAllFile()
	fmt.Println("查找的文件")
}
