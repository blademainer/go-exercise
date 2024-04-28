package main

import (
	"fmt"
	"reflect"
)

func F(af func() string) {

}

func main() {
	of := reflect.TypeOf(F)
	if of.Kind() == reflect.Func {
		for i := 0; i < of.NumIn(); i++ {
			fmt.Println("func in: ", of.In(i))
		}sdfkjalsokdf
	}
	fmt.Println(of.String())
}
