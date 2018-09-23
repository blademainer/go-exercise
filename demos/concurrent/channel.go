package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func(receiver chan<- int) {
		for i := 0; i < 100; i++ {
			receiver <- i
		}
		// compile error...
		//fmt.Println(<-receiver)
	}(c)

	func(sender <-chan int) {
		// produce: ./channel.go:18:10: invalid operation: sender <- 1 (send to receive-only type <-chan int)
		// sender <- 1
		for value := range sender {
			fmt.Println(value)
		}
	}(c)

	time.Sleep(1 * time.Second)
}
