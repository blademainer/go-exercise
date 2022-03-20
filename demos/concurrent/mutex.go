package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type A struct {
		name *string
		sync.Mutex
	}

	a := &A{}
	done := make(chan bool)
	for i := 0; i <= 100; i++ {
		go func(c chan bool) *string {
			time.Sleep(1 * time.Second)
			if a.name == nil {
				a.Lock()
				if a.name == nil {
					s := fmt.Sprintf("%d", i)
					a.name = &s
					fmt.Printf("name addr=== %s \n", a.name)
				}
				a.Unlock()
			}
			c <- true
			return a.name
		}(done)
	}
	for i := 0; i <= 100; i++ {
		<-done
	}
	fmt.Println(*a.name)
}
