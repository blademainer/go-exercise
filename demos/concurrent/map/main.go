package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(m map[string]string) {
		defer wg.Done()
		// fatal error: concurrent map iteration and map write
		for i := 0; i < 100; i++ {
			m[fmt.Sprintf("a%d", i)] = "b"
		}
	}(m)

	go func() {
		defer wg.Done()
		// fatal error: concurrent map iteration and map write
		for i := 0; i < 100; i++ {
			fmt.Println(m)
		}
	}()
	wg.Wait()
}
