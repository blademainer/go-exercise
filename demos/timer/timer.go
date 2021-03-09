package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	tk := time.NewTicker(time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case c := <-tk.C:
				fmt.Println("ticker, ", c.Format(time.RFC3339Nano))
			case <-ctx.Done():
				fmt.Println("done")
				return
			}
		}
	}()
	tm := time.NewTimer(time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case c := <-tm.C:
				fmt.Println("timer, ", c.Format(time.RFC3339Nano))
			case <-ctx.Done():
				fmt.Println("done")
				return
			}
		}
	}()

	wg.Wait()

}
