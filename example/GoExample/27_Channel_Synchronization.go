package main

import (
	"fmt"
	"time"
)

//通道同步，在最后添加 <-done 可实现worker函数的同步。
//如果去除最后的 <-done ，worker函数是执行不了的。

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done //只有在done接收到 ture 时程序才会往下走，而这只有等goroutine的worker函数自行完，
	//才会得到done <- true。所以这里充当阻塞的作用
	//而这也只是针对一个goroutine的同步方法
	//若是需要针对多个goroutine则是使用WaitGroup.
}
