package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	urls := []string{"https://www.sogou.com/tx?hdq=sogou-wsse-3f7bcd0b3ea82268-0001&ie=utf-8&query=", "https://www.sogou.com/?pid=sogou-wsse-3f7bcd0b3ea82268-0001"}
	for _, url := range urls {

		//resp, err := http.Get("https://www.sogou.com/?pid=sogou-wsse-3f7bcd0b3ea82268-0001")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error: %s \n", err)
			return
		}

		httpBody, _ := ioutil.ReadAll(resp.Body)
		body := string(httpBody)

		fmt.Println(body)
	}
}
