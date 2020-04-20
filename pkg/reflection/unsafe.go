package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type s struct {
	a bool
}

func main() {
	ts := &s{a: false}
	t := &s{a: true}
	fmt.Println(ts.a)
	p := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&ts)), unsafe.Pointer(ts), unsafe.Pointer(t))
	if !p {
		panic("failed to convert")
	}
	fmt.Println(ts.a)
}
