package main

import (
	"strings"
	"fmt"
)

func main() {
	//urls := []string{"\"https://news.qq.com/a/20180401/000109.htm\"", "\"https://news.qq.com/a/20180331/010784.htm\""}
	//for i, url := range urls {
	//	newString := strings.Replace(url, "\"", "", 0)
	//	urls[i] = newString
	//}
	//fmt.Println(urls)
	a := "\"asdgb\""
	fmt.Println(a)
	replace := strings.Replace(a, "\"", "", -1)
	fmt.Println(replace)
}
