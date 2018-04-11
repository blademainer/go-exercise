package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	dir := "html-out"
	name := "test.tmp"
	body := "hello!"
	_ = os.Mkdir(dir, os.ModePerm)
	path := dir+"/"+name
	file, e := os.Create(path)
	if e != nil{
		fmt.Println(e)
	}
	err := ioutil.WriteFile(path, []byte(body), 777)
	fmt.Println(err)
	file.Close()
}
