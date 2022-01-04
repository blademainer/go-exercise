package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	printSlice(newNilSlice())
	fmt.Println("----")
	printSlice(newEmptySlice())
}

func printSlice(ta []string) {
	marshal, err := json.Marshal(ta)
	if err != nil {
		panic(err.Error())
	}

	of := reflect.TypeOf(marshal)
	vf := reflect.ValueOf(ta)
	fmt.Printf(
		"value: %v type: %v value: %v nil: %v reflectNil: %v canAddr: %v \n", string(marshal), of,
		vf, ta == nil, vf.IsNil(),
		vf.CanAddr(),
	)
}

func newNilSlice() []string {
	return nil
}

func newEmptySlice() []string {
	return make([]string, 0)
}
