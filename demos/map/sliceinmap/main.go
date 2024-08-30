package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"a": make([]string, 0),
	}
	a := m["a"]
	fmt.Println(a[0])
}
