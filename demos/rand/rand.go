package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1)
	i := rand.Int()
	fmt.Println(i)

	rand.Seed(1)
	i2 := rand.Int()
	fmt.Println(i2)
}
