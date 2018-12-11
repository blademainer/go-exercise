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
	return not(or(and(a, b), and(not(a), not(b))))
}

func main() {
	fmt.Println(xor(true, false)) // true
	fmt.Println(xor(false, true)) // true
	fmt.Println(xor(true, true)) // false
	fmt.Println(xor(false, false)) // false
}
