package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s string
	err := json.Unmarshal([]byte("a string"), &s)
	if err != nil {
		fmt.Println("not a string", err.Error())
	}
	fmt.Println("s===", s)
	var s2 string
	err2 := json.Unmarshal([]byte(`"a2 string"`), &s2)
	if err2 != nil {
		fmt.Println("not a string", err2.Error())
	}
	fmt.Println("s2===", s2)
}
