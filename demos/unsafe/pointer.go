package main

import (
	"fmt"
	"unsafe"
)

func main() {
	h := 1
	p := &h
	var l = (*int64)(unsafe.Pointer(p))
	fmt.Println(l)
}
