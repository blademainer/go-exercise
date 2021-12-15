package main

import "fmt"

type a struct {
}

var emptyA = &a{}

var emptyV = a{}

func main() {
	emptyB := &a{}
	fmt.Println(emptyB == emptyA)  // false
	fmt.Println(*emptyB == emptyV) // true
}
