package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a struct {
		name string
		sync.Once
	}
	size := 10000
	done := make(chan bool)
	for i := 0; i < size; i++ {
		go func(exit chan bool, current int) {
			time.Sleep(time.Second)
			a.Do(
				func() {
					s := fmt.Sprintf("%s%d", "hello! ", current)
					fmt.Println("setting field name...", s)
					a.name = s
				},
			)
			fmt.Println("a.name: ", a.name)
			exit <- true
		}(done, i)
	}

	for i := 0; i < size; i++ {
		<-done
	}

}
