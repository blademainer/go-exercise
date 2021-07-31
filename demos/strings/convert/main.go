package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	b := []byte{0, 1, 2, 3, 255}
	s := string(b)
	fmt.Println(s)
	b2 := []byte(s)
	fmt.Println(len(b2))
	equal := reflect.DeepEqual(b, b2)
	if !equal {
		panic("not equal")
	}

	type a struct {
		Name string
	}

	aa := &a{
		Name: s,
	}
	fmt.Println(len([]byte(aa.Name)))
	marshal, _ := json.Marshal(aa)
	fmt.Println(string(marshal))
	aa2 := &a{}
	_ = json.Unmarshal(marshal, aa2)
	fmt.Println(aa2.Name)
	fmt.Println(len([]byte(aa2.Name)))
}
