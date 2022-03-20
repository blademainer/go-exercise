package main

import "fmt"

func main() {
	dataCh := make(chan int, 1024)
	fmt.Printf("len: %v cap: %v\n", len(dataCh), cap(dataCh))

	type a struct {
	}

	dc := make(chan *a, 1)
	dc <- nil
	fmt.Println(<-dc)
}
