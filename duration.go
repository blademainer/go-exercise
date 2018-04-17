package main

import (
	"time"
	"fmt"
)

func main() {
	seconds := 10
	duration := 10 * time.Second
	duration2 := time.Duration(time.Duration(seconds) * time.Second)
	fmt.Printf("type: %T value: %v", duration, duration)
	fmt.Printf("type: %T value: %v", duration2, duration2)
}
