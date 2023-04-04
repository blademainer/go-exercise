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

	aa := []int{1, 2, 4, 5}
	a3 := aa[:]
	fmt.Println(append(a3, 1))
	fmt.Println(a3)

	// go copy
	left := []int{1, 2, 3}
	right := []int{4, 5, 6, 7}
	all := make([]int, len(left)+len(right))
	copy(all[:len(left)], left)
	copy(all[len(left):], right)

	fmt.Println(all)
}
