package main

import (
	"strconv"
)

func main() {
	i := 0
	m := map[int]string{i: "b"}

	for range m {
		m[i] = strconv.Itoa(i)
		i++
	}
}
