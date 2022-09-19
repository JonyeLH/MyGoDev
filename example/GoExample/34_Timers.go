package main

import (
	"fmt"
	"time"
)

//时间器，在时间器等待时间结束后有管道接收数据
//时间器可以提前停止stop()
//Sheep()却不能中断
func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(2 * time.Second)
}
