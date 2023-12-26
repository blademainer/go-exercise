package main

import (
	"fmt"
)

func appendSlice(s []string) {
	s = append(s, "a")
}

func main() {
	var s []string
	appendSlice(s)
	fmt.Println(len(s))
}
