package main

import "fmt"

func main() {
	var a []byte
	fmt.Println(a == nil)
	fmt.Println(cap(a))
	fmt.Println(len(a))
	a = append(a, 1)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
