package main

import (
	"fmt"
	"reflect"
)

func main() {
	type a struct {

	}

	of := reflect.TypeOf(a{})
	fmt.Println("type: ", of)
}
