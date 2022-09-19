package main

import "fmt"

//channel管道，连接两个goroutine。可将管道信息传输给一个goroutine，然后将信息在另一个goroutine体现。
func main() {
	massage := make(chan string)

	go func() { massage <- "GOGOGO" }()

	MSG := <-massage
	fmt.Printf("massage chan to MSG %s\n", MSG)
}
