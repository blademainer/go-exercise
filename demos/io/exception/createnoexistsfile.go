package main

import (
	"os"
)

func main() {
	for {
		_, err := os.Create("tmp/test.log")
		if err == nil {
			panic("bad test")
		}
	}
}
