package main

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

func main() {
	number, err := decimal.NewFromString("20250811074413372133086")
	if err != nil {
		panic(err)
	}
	bigInt := number.BigInt()
	text := bigInt.Text(62)
	fmt.Println(text)

	ni := &big.Int{}
	_, ok := ni.SetString(text, 62)
	if !ok {
		panic("bad number: " + text)
	}
	fmt.Println(ni.String())
}
