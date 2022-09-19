package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nameList = map[string]string{"姓名": "李四", "性别": "男"}
	// 假如 key 存在，则 ok = true，否则，ok = false
	if name, ok := nameList["性别"]; ok {
		fmt.Println("Key存在，值为= ", name)
	} else {
		fmt.Println("Key不存在")
	}

	var b interface{}
	A, ok := b.(string)
	if ok == false {
		fmt.Println(reflect.TypeOf(A))
	}
}
