package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.DefaultClient.Get("http://unkonwn")
	if resp != nil && resp.Body != nil{
		defer resp.Body.Close()
	}
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

}
