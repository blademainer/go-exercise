package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	k int8  // 1字节
	i int32 // 4字节
	j int64 // 8字节
}

func main() {
	var t = &A{}
	var a = (*int8)(unsafe.Pointer(t))
	*a = 1
	var b = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(t)) + unsafe.Offsetof(t.i))) // Offsetof(x.y) y字段相对于x起始地址的偏移量，包括可能的空洞。
	*b = 2
	var c = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(t)) + unsafe.Offsetof(t.j)))
	*c = 3
	fmt.Printf("%#v\n", t)
}

// 输出：&main.A{k:1, i:2, j:3}
