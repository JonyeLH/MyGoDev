package main

import "fmt"

//关闭管道
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receiced job", j)
			} else { //只有在完成所有jobs的数据接收 且 关闭jobs管道 进入else判断
				fmt.Println("finish received jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")
	<-done
}
