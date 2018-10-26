package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go func() {
		//defer func() {
		//	if i := recover(); i != nil{
		//		fmt.Println("Panic with: ", i)
		//	}
		//}()
		resp, err := http.DefaultClient.Get("http://unkonwn")
		resp.Body.Close()
		//if resp != nil && resp.Body != nil{
		//	defer resp.Body.Close()
		//}
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("aaaa")
}
