package main

import (
	"fmt"
	"math/rand"
	"sync/atomic" //原子操作包atomic
	"time"
)

type Readop struct {
	key  int
	resp chan int
}

type Writeop struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readop uint64
	var writeop uint64

	reads := make(chan Readop)
	writes := make(chan Writeop)

	//区分读写操作，采用map完成数据结构
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//通过goroutine和channel，完成数据的读写的管控
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := Readop{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readop, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := Writeop{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeop, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readopfinal := atomic.LoadUint64(&readop) //atomic.LoadUint64(*)是原子操作的加载，加载不再改变的值
	fmt.Println("readopfinal:", readopfinal)
	writeopfinal := atomic.LoadUint64(&writeop)
	fmt.Println("writeopfinal:", writeopfinal)

}
