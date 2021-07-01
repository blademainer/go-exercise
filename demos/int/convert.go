package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("string cast: ", string(199))
	fmt.Println("string cast: ", string(1))
	fmt.Println("strconv: ", strconv.FormatInt(199, 10))
}
