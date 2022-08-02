package main

import (
	"fmt"
)

type A struct {
	b *B
}

type B struct {
	Name string
	c    *C
}

type C struct {
	CName string
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
	a := &A{
		b: &B{
			Name: "init",
			c:    &C{CName: "initc"},
		},
	}

	fmt.Println(a.b.Name)
	fmt.Println(a.b.c.CName)

	// bad clone
	na := a.BadClone()
	na.b.Name = "new"
	na.b.c.CName = "newc"
	fmt.Println(na.b.Name)   // new
	fmt.Println(a.b.Name)    // new
	fmt.Println(a.b.c.CName) // newc

	// good clone
	na2 := a.Clone()
	na2.b.Name = "good"
	na.b.c.CName = "goodc"
	fmt.Println(na2.b.Name) // good
	fmt.Println(a.b.Name)   // new
	fmt.Println(a.b.c.CName)

}
