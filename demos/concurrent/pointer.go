package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

var ptr *unsafe.Pointer

type a struct {
	name string
}

func main() {
	p := unsafe.Pointer(nil)
	ptr = &p
	fmt.Println(atomic.LoadPointer(ptr))
	atomic.CompareAndSwapPointer(ptr, unsafe.Pointer(nil), unsafe.Pointer(&a{"zhangsan"}))
	fmt.Println((*a)(atomic.LoadPointer(ptr)))
}
