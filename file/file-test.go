package main

import (
	"fmt"
	"io/ioutil"
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
