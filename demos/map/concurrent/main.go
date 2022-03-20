package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	i := 0
	go func() {
		// fatal error: concurrent map read and map write
		// couldn't recover
		defer func() {
			fmt.Println(recover())
		}()
		for {
			m[i] = i
			i++
		}
	}()
	go func() {
		defer func() {
			fmt.Println(recover())
		}()
		for {
			_ = m[i]
		}
	}()

	select {}
}
