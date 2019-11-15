package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	once := sync.Once{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			time.Sleep(10 * time.Millisecond)
			once.Do(func() {
				fmt.Println("init!!!", index) // random!!!
			})
			wg.Done()
		}(i)
	}

	wg.Wait()
}
