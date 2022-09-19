package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/defer.txt")
	defer closeFile(f) //defer 在main函数关闭之前执行
	writeFile(f)
}

func createFile(str string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(str)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
