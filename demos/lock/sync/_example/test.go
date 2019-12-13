package main

import (
	"fmt"
	locksync "github.com/blademainer/go-exercise/demos/lock/sync"
	"sync"
)

func main() {
	a := make(map[int]int)
	rw := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			sum := i + j
			k := sum % 10
			go func() {
				exists := -1
				locksync.Do(rw,
					func() bool {
						var found bool
						exists, found = a[k]
						return !found
					},
					func() {
						fmt.Printf("found: %v replace with: %v\n", exists, sum)
						a[k] = sum
					},
				)
				wg.Done()
			}()

		}
	}
	wg.Wait()
}
