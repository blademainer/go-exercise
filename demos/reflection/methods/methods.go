package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
}

func (a *A) String() string {
	return "A"
}

func (a *A) Hello() string {
	return "Hello " + a.Name
}

func AOPOfStruct(p any) {
	of := reflect.ValueOf(p)
	fmt.Printf("of.Type(): %v of.NmMethod: %v\n", of.Type(), of.NumMethod())
	nf := of.Elem()
	for i := 0; i < nf.NumMethod(); i++ {
		method := nf.Method(i)
		if !method.CanAddr() {
			panic("method cannot addr")
		}
		AOPOfFunc(method) // 获取method的地址
	}
}

func AOPOfFunc(f reflect.Value) {
	// Copy is needed in order to prevent infinite recursion after function wrapping.
	fmt.Println("f.Type():", f.Type())
	oldDoValueCopy := reflect.New(f.Type()).Elem()
	oldDoValueCopy.Set(f)
	fv := reflect.MakeFunc(f.Type(), func(args []reflect.Value) []reflect.Value {
		fmt.Println("before")
		fmt.Printf("args: %v\n", args)
		result := oldDoValueCopy.Call(args)
		fmt.Println("after")
		fmt.Printf("result: %v\n", result)
		return result
	})
	f.Set(fv)
}

func main() {
	a := &A{Name: "world"}
	AOPOfStruct(&a)
	fmt.Println(a.Hello())
	fmt.Println(a.String())
}
