package main

import (
	"fmt"
	"sync"
)

type A struct {
	sync.Mutex
}

// Test test
func Test(a A) {

}

func main() {
	m := make(map[int]sync.Once)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			once := m[i%10] // Variable declaration copies lock value to 'once': type 'sync.Once' contains 'sync.Mutex' which is 'sync.Locker'
			once.Do(
				func() {
					fmt.Println("init: ", i%10)
				},
			)
		}()
	}
	wg.Wait()
}
