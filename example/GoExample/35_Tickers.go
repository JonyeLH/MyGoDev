package main

import (
	"fmt"
	"time"
)

func main() {
	tickers := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tickers.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	tickers.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
