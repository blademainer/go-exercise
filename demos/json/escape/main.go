package main

import (
	"bytes"
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

	// disable EscapeHTML
	bts := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(bts)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(aa)
	if err != nil {
		panic(err)
	}
	fmt.Println(bts.String())

	// unmarshal
	err = json.Unmarshal(bts.Bytes(), &aa)
	if err != nil {
		panic(err)
	}
	fmt.Println(aa)
}
