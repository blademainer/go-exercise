package main

import (
	"fmt"
)

func main() {
	fmt.Printf("app\rcd\n") // cd
	fmt.Println("---")
	// app
	// cd
	fmt.Printf("app\r\ncd")
	fmt.Println("---")
	fmt.Printf("app\n\rcd")
	fmt.Println("---")
}
