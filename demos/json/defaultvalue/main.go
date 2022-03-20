package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type a struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}
	type b struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	marshal, _ := json.Marshal(&a{})
	fmt.Println(string(marshal)) // {}
	marshal, _ = json.Marshal(&b{})
	fmt.Println(string(marshal)) // {"name":"","age":0}
}
