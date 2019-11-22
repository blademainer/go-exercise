package main

import "fmt"

// !a
func not(a bool) bool {
	return !a
}

// a & b
func and(a, b bool) bool {
	return a && b
}

// a | b
func or(a, b bool) bool {
	return a || b
}

// a ^ b
func xor(a, b bool) bool {
	return and(or(a, b), or(not(a), not(b))) // (!a|!b)&(a|b)
}

func main() {
	fmt.Println(xor(true, false))  // true
	fmt.Println(xor(false, true))  // true
	fmt.Println(xor(true, true))   // false
	fmt.Println(xor(false, false)) // false
}
