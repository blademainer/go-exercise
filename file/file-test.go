package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	infos, e := ioutil.ReadDir("./")
	if e != nil {
		fmt.Println("error === ", e.Error())
		return
	}

	for _, info := range infos {
		fmt.Println("file==== ", info)
	}
}
