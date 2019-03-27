package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func request(context context.Context, data chan int) {
	defer wg.Done()
	for {
		select {
		case i := <-data:
			fmt.Println("receive data: ", i)
		case <-context.Done():
			fmt.Println("Done...")
			return
		}
	}
}

func main() {
	ints := make(chan int, 1024)
	go func() {
		i := 0
		for {
			ints <- i
			i++
		}
	}()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	wg.Add(1)
	go request(ctx, ints)
	wg.Wait()
}
