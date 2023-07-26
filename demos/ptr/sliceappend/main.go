package main

import (
	"fmt"
	"strconv"
)

type a struct {
	b string
}

func (a *a) String() string {
	return a.b
}

func main() {
	var aa *a
	var as []*a
	for i := 0; i < 10; i++ {
		aa = &a{b: strconv.Itoa(i)}
		as = append(as, aa)
	}
	fmt.Println(as)
}
