package main

import (
	"github.com/go-exercise/util/collection"
	"fmt"
)

func MultiInsert(){

}

func main() {
	list := collection.Init()
	for i := 0; i < 100000; i++{
		list.Add(i)
	}
	fmt.Printf("%T \n", list.First())
	fmt.Println(list.Last())
	fmt.Println(list.Size)

}
