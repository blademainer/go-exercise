package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Started goroutine...")
		panic("aaaa")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("OK?")
}
