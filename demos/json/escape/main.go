package main

import (
	"encoding/json"
	"fmt"
)


func main() {
	type a struct {
		URL string `json:"url"`
	}

	aa := &a{URL: "https://test.com?name=zhangsan&age=123&email=test@gmail.com"}

	raw, err := json.Marshal(aa)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))

	err = json.Unmarshal(raw, &aa)
	if err != nil {
		panic(err)
	}
	fmt.Println(aa)
}
