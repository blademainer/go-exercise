package main

import (
	"time"
	"fmt"
	"sync/atomic"
)

func main() {
	bools := make(chan int)
	var count int32 = 1
	go func() {
		for {
			fmt.Println("count: ", count)
		}
	}()
	for i := 0; i < 1000000; i++ {
		go func(done chan int, index int) {
			atomic.AddInt32(&count, 1)
			time.Sleep(10 * time.Second)
			//fmt.Println("index: ", index)
			done <- index
			atomic.AddInt32(&count, -1)
		}(bools, i)
	}

	for i := 0; i < 1000000; i++ {
		<-bools
		//fmt.Println("done: ", index)
	}
}
