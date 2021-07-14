package main

import (
	"errors"
	"fmt"
)

var (
	errTmp = errors.New("tmp err")
)

type errType string

func (e errType) Error() string {
	return fmt.Sprintf(string(e))
}

func (e errType) Validate() bool {
	return true
}

func main() {
	fmt.Println(errors.Is(errTmp, errTmp))
	fmt.Println(errors.Is(errType("test err"), errType("test err")))
	fmt.Println(errors.Is(errType("test err"), errType("test err2")))
	fmt.Println(errTmp.Error())
	fmt.Println(errType("test err").Error())
	fmt.Println(Validate(errType("test err")))
}

func Validate(i interface{}) bool {
	if v, ok := i.(interface {
		Validate() bool
	}); ok {
		return v.Validate()
	}
	return false
}
