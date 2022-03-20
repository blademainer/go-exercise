package main

import (
	"encoding/json"
	"fmt"
)

type In struct {
	Code    int    `json:"code"`
	CodePtr *int   `json:"code_ptr"`
	Message string `json:"message"`
}

func prin(i *In) {
	bytes, _ := json.Marshal(i)
	fmt.Println(string(bytes))

	in := &In{}
	json.Unmarshal(bytes, in)
	fmt.Println(in)
}

func main() {
	in := &In{}
	prin(in)

	in2 := &In{}
	in2.Code = 0
	in2.CodePtr = &in2.Code
	in2.Message = "SUCCESS"
	prin(in2)

}
