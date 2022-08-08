package main

import (
	"fmt"
	"strconv"
)

func main() {
	amount, err := strconv.ParseFloat("1212345678.129123123", 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f", amount)
}
