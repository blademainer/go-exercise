package main

import "fmt"

type People struct {
	name string
	age  int8
}

func Test(p interface{}) {
	pe := p.(People)
	fmt.Println(pe)
}

func main() {
	Test(People{"a", 1})
	Test("")
}
