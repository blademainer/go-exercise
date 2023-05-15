package generic

import (
	"fmt"
	"reflect"
)

func ReflectType(t interface{}) {
	of := reflect.TypeOf(t)
	fmt.Println(of)
	elem := of.Elem()
	fmt.Println(elem)
	fmt.Println(elem)
}

func NewT[T any]() *T {
	tt := new(T) // good
	// tt := T{} // bad
	// tt := &T{} // bad
	return tt
}
