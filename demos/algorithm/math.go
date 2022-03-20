package main

import (
	"fmt"
	"math"
)

func xx(a int) (result int) {
	x := float64(a)
	f := (19053.0+(5.0/6.0))*x*x*x - 114263.0*x*x + (209290.0+(1.0/6.0))*x - 113244
	result = int(math.Round(float64(f)))
	return result
}

func main() {
	fmt.Println(xx(1))
	fmt.Println(xx(2))
	fmt.Println(xx(3))
	fmt.Println(xx(4))
}
