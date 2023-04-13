package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func writeSlice(p interface{}) {
	of := reflect.ValueOf(p)
	pv := of.Elem()

	newv := pv
	if pv.Cap() < 1 {
		newv = reflect.MakeSlice(pv.Type(), 0, 1)
		reflect.Copy(newv, pv)
	}

	// elem := of.Elem()
	et := pv.Type().Elem()
	fmt.Println("et: ", et)
	value := reflect.New(et)
	fmt.Println("value: ", value.Type())
	fmt.Println(value.Elem())
	fmt.Println(value.Interface())
	fmt.Println(value.Type().Elem())
	err := json.Unmarshal([]byte(`{"Name": "zhangsan"}`), value.Interface())
	if err != nil {
		panic(err)
		return
	}
	newv = reflect.Append(newv, value.Elem())
	// pv.Index(0).Set(value.Elem())
	pv.Set(newv)

	fmt.Printf("%#v\n", value.Elem().Interface())
}

func main() {
	type p struct {
		Name string
	}
	var pp []*p
	writeSlice(&pp)
	fmt.Println("pp len: ", len(pp))
	for _, p2 := range pp {
		fmt.Printf("p2: %#v\n", p2)
	}
	var pp2 []p
	writeSlice(&pp2)
	for _, p2 := range pp2 {
		fmt.Printf("_p2: %#v\n", p2)
	}
}
