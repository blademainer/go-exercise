package main

import (
	"fmt"
	"net/url"
)

func main() {
	query, err := url.ParseQuery("a=1&b=2&c=")
	if err != nil {
		return
	}
	c, ok := query["c"]
	if !ok {
		fmt.Println("c is not exist")
	}
	fmt.Println(len(c))
	fmt.Println(c[0])
}
