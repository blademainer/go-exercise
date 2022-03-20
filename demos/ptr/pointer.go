package main

import "fmt"

func t(a *[]string) {
	*a = append(*a, "1")
}
func c(a []string) {
	a[0] = "b"
}

func main() {
	a := make([]string, 0, 10)
	fmt.Println(len(a), cap(a))
	t(&a)
	t(&a)
	fmt.Println(len(a))
	a[0] = "a"
	c(a)
	fmt.Println(a[0])
}
