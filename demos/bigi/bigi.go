package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := int64(1)
	bi := big.NewInt(i)
	leftMove(bi, 1000000)
	fmt.Println(bi)
}

func leftMove(bi *big.Int, size int) {
	for i := 0; i < size; i++ {
		*bi = *(big.NewInt(0).Mul(bi, big.NewInt(2)))
	}
}
