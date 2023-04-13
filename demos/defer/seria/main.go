package main

import (
	"fmt"
)

func testDefer() {
	defer fmt.Println("d1")
	defer fmt.Println("d2")
	fmt.Println("d3")
}

func main() {
	testDefer()
}
