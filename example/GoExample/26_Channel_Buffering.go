package main

import "fmt"

//通道缓冲区，在声明通道的时候使用第二参数，表示通道的缓冲大小。
//不添加第二个参数，表示没有缓冲数据
func main() {
	//作为接收数据的管道要设定其管道容量
	//作为传输数据的管道则不需要设定容量
	messages := make(chan string, 2) //通道缓冲区大小为2

	messages <- "buffered" //可以同时进行若干个通道传输数据，而不用立刻需要去同步读取数据
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
