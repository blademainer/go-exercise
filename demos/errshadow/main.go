package main

import (
	"fmt"
	"strconv"
)

func foo(x string) (ret int, err error) {
	if true {
		_, err := strconv.Atoi(x)
		if err != nil {
			// should compile error
			return
		}
	}
	return ret, nil
}

func main() {
	fmt.Println(foo("123"))
}
