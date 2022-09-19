package main

import "fmt"

//通道连接，将两个管道的信息连通。
//将pings管道的信息连接到pongs，并输出表示。

//ping函数定义的输入参数  pings chan<- string  表示pings这个管道是作为接收数据。
//根据<-指向的内容确定，哪个作为接收。
func ping(pings chan<- string, str string) {
	pings <- str
}

//在pong函数定义的输入参数  pings <-chan string  表示pings管道作为数据传输
//  pongs chan<- string  表示pongs管道作为接收数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	//作为接收数据的管道要设定其管道容量
	//作为传输数据的管道则不需要设定容量
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed massage!")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
