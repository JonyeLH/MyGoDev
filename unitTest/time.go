package main

import (
	"fmt"
	"time"
)

func TimeStamp() {
	timeNow := time.Now()
	time.Sleep(5 * time.Second)
	TimeDuration := time.Now().Sub(timeNow).Seconds()
	fmt.Println("timestamp = ", TimeDuration)
}
