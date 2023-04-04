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
}
