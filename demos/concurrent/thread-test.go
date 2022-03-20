package main

import (
	"fmt"

	"github.com/blademainer/go-exercise/util/collection"
)

func add(list *collection.List[int], i chan int) {
	for j := 0; j <= 100; j++ {
		i <- j
		list.Add(j)
	}
	close(i)
}

func main() {
	i := make(chan int)
	list := collection.New[int]()
	go add(list, i)
	go add(list, i)
	go add(list, i)
	for n := range i {
		fmt.Println(n)
	}
	fmt.Println(list.Size)
}
