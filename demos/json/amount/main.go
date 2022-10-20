package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Amount json.Number
}

func main() {
	o := &Order{
		Amount: json.Number("123.45"),
	}
	marshal, err := json.Marshal(o)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}
