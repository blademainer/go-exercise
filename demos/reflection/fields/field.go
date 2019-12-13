package main

import (
	"fmt"
	"reflect"
)

type Root struct {
	Child *Child
}

type Child struct {
	Name string
	Age  int
}

func main() {
	c := &Child{Name: "zhangsan"}
	r := &Root{Child: c}
	of := reflect.ValueOf(r).Elem()
	fmt.Println(of.Field(0))
}
