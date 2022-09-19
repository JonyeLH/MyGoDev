package main

import (
	"fmt"
	"sync"
	"time"
)

func workerss(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Println("sleep a second")
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup //声明WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) //WaitGroup计数器，用于标识
		i := i    //避免闭包中goroutine的i复用

		go func() {
			defer wg.Done() //将worker包括在闭包中，使用defer确保worker执行完毕
			workerss(i)
		}()
	}
	wg.Wait() //阻塞到WaitGroup的计数器为0，表示所有的worker完成
}
