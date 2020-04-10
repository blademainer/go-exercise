package main

import (
	"fmt"
	"reflect"
)

func main() {
	type t struct {
		name string
	}

	tt := &t{name: "zhangsan"}
	of := reflect.TypeOf(tt)
	value := reflect.New(of.Elem())
	t2 := value.Interface().(*t)
	t2.name = "clone of zhangsan"
	fmt.Println(t2)
}
