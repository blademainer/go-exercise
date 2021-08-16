package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	type a struct {
		Age  int
		F    float32
		Name string
		Age2 int
	}

	var (
		req = a{}
		rsp = &a{}
	)

	wg := &sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		i := i * i * i
		go func() {
			defer wg.Done()
			for {
				rsp.Name = strconv.Itoa(i)
				fmt.Sprintf("req and rsp \n", req, *rsp)
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				rsp.Age = i
				rsp.Age2 = i
				rsp.F = float32(i)
				fmt.Sprintf("req and rsp \n", req, *rsp)
			}
		}()
	}
	wg.Wait()

}
