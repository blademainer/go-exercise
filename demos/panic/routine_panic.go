package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Started goroutine...")
		panic("aaaa")
	}()
	time.Sleep(2 * time.Second)
}
