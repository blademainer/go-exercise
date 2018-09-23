package main

import "fmt"

func main() {
	a := "hello"
	b := "world"
	*(&a), *(&b) = b, a

	fmt.Printf("a=%v, b=%v \n", a, b)
	c := &a
	d := &c

	fmt.Println("aaa")
	fmt.Printf("Type=%T \n", d)
}

func pp(in interface{}) interface{} {
	for i := 0; i <= 100; i++ {
		in = &in
		fmt.Printf("Type=%T \n", in)
	}
	return in
}
