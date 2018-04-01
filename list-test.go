package main

import (
	"github.com/go-exercise/util/collection"
	"fmt"
)

func MultiInsert() {

}

func main() {
	list := collection.Init()
	for i := 0; i < 100000; i++ {
		list.Add(i)
	}
	fmt.Printf("First type: [%T]  value: %v \n", list.First(), list.First())
	fmt.Println("Last: ", list.Last())
	fmt.Println("Size: ", list.Size)

}
