package main

import (
	"fmt"
)

func main() {
	type a struct {
		names []string
	}

	aa := &a{}
	aa.names = append(aa.names, "test")
	fmt.Println(aa)
}
