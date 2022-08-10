package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make([]string, 0)
	m = append(m, "1")
	for i, v := range m {
		m = append(m, strconv.Itoa(i))
		fmt.Println(v)
	}
}
