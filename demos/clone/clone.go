package main

import (
	"fmt"
)

type A struct {
	b *B
}

type B struct {
	Name string
}

func (a *A) BadClone() *A {
	na := *a
	return &na
}

func (a *A) Clone() *A {
	na := *a
	nb := *(a.b)
	na.b = &nb
	return &na
}

func main() {
	a := &A{b: &B{Name: "init"}}
	fmt.Println(a.b.Name)

	// bad clone
	na := a.BadClone()
	na.b.Name = "new"
	fmt.Println(na.b.Name) // new
	fmt.Println(a.b.Name)  // new

	// good clone
	na2 := a.Clone()
	na2.b.Name = "good"
	fmt.Println(na2.b.Name) // good
	fmt.Println(a.b.Name)   // new
}
