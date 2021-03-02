package main

import "fmt"

type A string

type SliceKey struct {
	slice []A
	string
}

func (s SliceKey) String() string {
	return fmt.Sprint(s.string)
}


func main() {
	strings := make(map[A]string)
	strings["123"] = "123"
	fmt.Println(strings)
}
