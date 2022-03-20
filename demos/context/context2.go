package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 = sync.WaitGroup{}

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
	wg2.Add(1)
	go requestTimeout(ctx, ints)
	wg2.Wait()
}

func requestTimeout(ctx context.Context, data chan int) {
	defer wg2.Done()

	timeout, cacelFunc := context.WithTimeout(ctx, 2*time.Second)
	for {
		select {
		case i := <-data:
			fmt.Println("receive data: ", i)
		case <-ctx.Done():
			fmt.Println("time out of parent... invoke child cacleFunc...")
			cacelFunc()
			return
		case <-timeout.Done():
			fmt.Println("time out of child...")
			return
		}
	}
}
