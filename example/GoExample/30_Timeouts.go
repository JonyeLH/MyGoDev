package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result1"
	}()
	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1s")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res) //在2秒时延后打印c2
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3s") //超过3秒打印时延
	}
}
