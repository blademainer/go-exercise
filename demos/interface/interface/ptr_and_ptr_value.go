package main

import (
	"fmt"
)

type I interface {
	SayHello()
}

type a struct {
	name string
}

func (a a) SayHello() {
	fmt.Println("hello value type, name: ", a.name)
}

type b struct {
	name string
}

func (b *b) SayHello() {
	fmt.Println("hello ptr type, name: ", b.name)
}

func SayHello(i I) {
	switch it := i.(type) {
	case a:
		fmt.Println("receive a value type")
		it.name = "say hello value motified"
	case *a:
		it.name = "say hello ptr motified"
	case *b:
		it.name = "say hello motified"
	}
	i.SayHello()
}

func main() {
	pa := &a{
		name: "default",
	}
	SayHello(pa)
	pa.name = "present"
	SayHello(pa)

	va := a{
		name: "default",
	}
	SayHello(va)
	va.name = "present value"
	SayHello(va)

	pb := &b{
		name: "default",
	}
	SayHello(pb)
}
