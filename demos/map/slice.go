package main

import "fmt"

type A string

type SliceKey struct {
	slice []A
	string
}

func (s SliceKey) String() string {
	return fmt.Sprint(s)
}


func main() {
	strings := make(map[SliceKey]string)
	strings[SliceKey{string: "123"}] = "123"
	fmt.Println(strings)
}
