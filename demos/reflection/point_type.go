package main

import (
	"fmt"
	"reflect"
)

func IsPointType(i interface{}) bool {
	of := reflect.TypeOf(i)
	if of.Kind() == reflect.Ptr {
		fmt.Printf("Type: %v is point type! \n", of)
		return true
	} else {
		fmt.Printf("Type: %v is not point type! \n", of)
		return false
	}
}

func main() {
	a := ""
	IsPointType(a)
	IsPointType(&a)
}
