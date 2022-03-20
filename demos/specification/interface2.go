package main

import "fmt"

type I2 interface {
	Set(int)
}

type T2 struct {
	i int
}

func (t2 *T2) Set(i int) {
	t2.i = i
}

func f(i I2) {
	i.Set(10)
}

func main() {
	t2 := &T2{}
	f(t2)
	fmt.Println(t2.i)
}
