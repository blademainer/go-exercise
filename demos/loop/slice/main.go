package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	s := make([]int, 1, 10)

	for range s {
		s = append(s, i)
		fmt.Println("range ", i)
		i++
	}

	for i := 0; i < len(s); i++ {
		fmt.Println("loop ", i)
		s = append(s, i)
		time.Sleep(1 * time.Second)
	}
}
