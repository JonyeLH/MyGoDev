package main

import "os"

//panic：得到不符合的结果，快速输出判错，无需执行后续代码
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
