package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ar := []string{"1", "2"}
	var hdr reflect.StringHeader
	hdr.Data = uintptr(unsafe.Pointer(&ar))
	hdr.Len = 2
	s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lostunsafe.Pointer(ar)
	fmt.Println(s)
}
