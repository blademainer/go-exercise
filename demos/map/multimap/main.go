package main

import (
	"fmt"
)

func main() {
	mm := make(map[string]map[string]string)
	mm["a"] = make(map[string]string)

	m := make(map[string]map[string]string)
	data, ok := m["a"]
	if !ok {
		data = make(map[string]string)
	}
	data["b"] = "1"
	m["a"] = data
	fmt.Println(m)

	m["c"]["d"] = "1" // panic
	fmt.Println(m)
}
