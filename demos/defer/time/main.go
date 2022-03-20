package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	defer fmt.Printf("test\n")
	defer fmt.Printf("%d\n", time.Since(a).Microseconds())
	defer func() {
		fmt.Printf("%d\n", time.Since(a).Microseconds())
	}()
	time.Sleep(3 * time.Second)
}
