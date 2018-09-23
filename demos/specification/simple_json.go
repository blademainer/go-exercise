package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	b := true
	bytes, _ := json.Marshal(b)
	res := string(bytes)
	fmt.Println(res)
	//	try convert
	a, _ := strconv.ParseBool(res)
	fmt.Println(a)

}
