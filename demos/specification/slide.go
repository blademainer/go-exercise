package main

import (
	"fmt"
)

func main() {
	bytes := []byte{1, 2, 3, 4, 5}
	fmt.Println(bytes)
	fmt.Println([]byte("\n"))
	fmt.Println(InsertNewLine(bytes))
}

func InsertNewLine(b []byte) []byte {
	target := make([]byte, len(b)+1)
	copy(target, []byte("\n"))
	copy(target[1:], b)
	return target
}
