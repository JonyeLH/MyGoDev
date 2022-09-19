package main

//使用管道、tick、goroutines完成接收数据的控制
import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)
	limiter := time.Tick(200 * time.Millisecond)
	// limiter := time.Tick(2 * time.Second)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//若是又突发数据需要接受，则再开含tick的管道来接收
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(3 * time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	//这里接收的前三个数据是通过突发管道及时接收的，无需等待
	//后两个需要等待tick时间
	for busReq := range burstyRequest {
		<-burstyLimiter
		fmt.Println("burstyRequest", busReq, time.Now())
	}
}
