package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func sum(number uint64) uint64 {
	if number <= 1 {
		return uint64(number)
	}

	return sum(number-1) + sum(number-2)
}

func main() {
	go func() {
		// open: http://localhost:6060/debug/pprof/
		result := http.ListenAndServe("localhost:6060", nil)
		fmt.Println("init result: ", result)
	}()
	for i := 0; i < 1000; i++ {
		result := sum(uint64(i))
		fmt.Printf("Fibonacci: %d result: %d \n", i, result)
	}
}
