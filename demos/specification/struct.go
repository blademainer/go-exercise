package main

import (
	"bytes"
	"fmt"
)

type State struct {
	length uint32
	bytes.Buffer
}

func main() {
	s := &State{}
	s.WriteByte('a')
	fmt.Println(string(s.Bytes()))
}
