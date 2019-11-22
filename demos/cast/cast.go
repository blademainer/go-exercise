package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]interface{})
	m["a"] = "b"
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		cc := m["cc"].(string) // panic
		fmt.Println(cc)
	}()
	wg.Wait()
	fmt.Println("done")
}
