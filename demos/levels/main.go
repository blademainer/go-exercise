package main

import (
	"fmt"
)

func main() {
	groupSize := 0
	_, err := fmt.Scanf("%d", &groupSize)
	if err != nil {
		panic(err)
	}
	for i := 0; i < groupSize; i++ {
		var m, s, c, l int
		_, err := fmt.Scanf("%d %d %d", &m, &s, &c)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Scanf("%d", &l)
		if err != nil {
			panic(err)
		}
		fmt.Print(minIcons(m, s, c, l))
		if i != groupSize-1 {
			fmt.Println()
		}
	}
}

func minIcons(m int, s int, c int, l int) int {
	sum := 0

	for l >= 0 {
		if l >= c*m*s {
			sum += l / (c * m * s)
			l = l % (c * m * s)
		} else if l >= s*m {
			sum += l / (s * m)
			l = l % (s * m)
		} else if l >= m {
			sum += l / m
			l = l % m
		} else {
			break
		}
	}

	if l > 0 {
		sum += l
	}
	return sum
}
