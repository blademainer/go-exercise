package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func process() func(*sync.WaitGroup, []string) error {
	return func(wg *sync.WaitGroup, strings []string) error {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			// loop strings
			for _, s := range strings {
				fmt.Println(s)
			}
		}()
		return nil
	}
}

func main() {
	wg := &sync.WaitGroup{}
	f := process()
	s := make([]string, 0)
	for i := 0; i < 100; i++ {
		s = make([]string, 0)
		s = append(s, "hello"+strconv.Itoa(i))
		err := f(wg, s)
		if err != nil {
			panic(err)
		}
	}
	wg.Wait()
}
