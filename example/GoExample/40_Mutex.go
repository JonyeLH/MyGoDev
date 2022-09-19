package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu        sync.Mutex
	container map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.container[name]++
}

func main() {
	c := Container{
		container: map[string]int{"a": 1000, "b": 0},
	}
	var wg sync.WaitGroup
	increase := func(name string, num int) {
		for i := 0; i < num; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go increase("a", 1000)
	go increase("b", 1000)
	go increase("a", 1000)
	wg.Wait()
	fmt.Println(c.container)
}
