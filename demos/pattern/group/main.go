package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(`\$\{(.*?)}`)
	submatch := pattern.FindAllStringSubmatch("${asd}", -1)
	fmt.Println(submatch[0][1]) // asd
}
