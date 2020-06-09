package main

import "fmt"

func main() {
	dataCh := make(chan int, 1024)
	fmt.Printf("len: %v cap: %v", len(dataCh), cap(dataCh))
}
