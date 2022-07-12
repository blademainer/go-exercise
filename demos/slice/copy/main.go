package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4, 5}

	// bad
	var na []int
	na = append(a[:2], 3)
	na = append(na, a[2:]...)
	fmt.Println(na)
	fmt.Println(a)

	// good
	a2 := []int{1, 2, 4, 5}
	var na2 []int
	na2 = append(na2, a2[:2]...)
	na2 = append(na2, 3)
	na2 = append(na2, a2[2:]...)
	fmt.Println(na2)
	fmt.Println(a2)

}
