package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 如果不设置种子，每次运行程序生成的随机数都是一样的
	rand.Seed(1)
	i := rand.Int()
	fmt.Println(i)

	time.Sleep(2 * time.Second)
	// 如果设置的种子一样，那么每次运行程序生成的随机数都是一样的
	rand.Seed(1)
	i2 := rand.Int()
	fmt.Println(i2)
}
