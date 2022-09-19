package main

import (
	"fmt"
	"time"
)

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	f("Holle")
	go f("World")

	go func(msg string) {
		fmt.Printf("Msg is: %s\n", msg)
	}("Go")

	time.Sleep(time.Second)
	fmt.Println("End!!!")
}
