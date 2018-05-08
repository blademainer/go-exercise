package main

import "fmt"

func main() {
	a := "hello"
	b := "world"
	*(&a), *(&b) = b, a

	fmt.Printf("a=%v, b=%v", a, b)
}
