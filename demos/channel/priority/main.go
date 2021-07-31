package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 1024)
	c2 := make(chan int, 1024)
	fmt.Println(cap(c1))
	fmt.Println(len(c1))
	fmt.Println(cap(c2))
	fmt.Println(len(c2))
	index := 0
	for i := 0; i < 1024; i++ {
		c1 <- index
		index++
	}
	for i := 0; i < 1024; i++ {
		c2 <- index
		index++
	}
	close(c1)
	close(c2)
	fmt.Println(cap(c1))
	fmt.Println(len(c1))
	fmt.Println(cap(c2))
	fmt.Println(len(c2))

	data := make([]int, 0, len(c1)+len(c2))
	for {
		var shutdown1 bool
		var shutdown2 bool

		select {
		case i, ok := <-c1:
			if !ok {
				shutdown1 = true
			} else {
				data = append(data, i)
				continue
			}
		default:
		}
		select {
		case i, ok2 := <-c2:
			if !ok2 {
				shutdown2 = true
			} else {
				data = append(data, i)
			}
		}

		if shutdown2 && shutdown1 {
			break
		}
	}
	for _, datum := range data {
		fmt.Println(datum)
	}
	fmt.Println(len(data))
}
