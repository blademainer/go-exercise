package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%.10f\n", float64(311232321.3232321321)+float64(311232321.3232321321))
	fmt.Printf("%.10f\n", float64(123123311232321.3232321321)+float64(123311232321.3232321321))
	amount, err := strconv.ParseFloat("1212345678.129123123", 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f\n", amount)
	amount, err = strconv.ParseFloat("16.123", 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f\n", amount)
}
