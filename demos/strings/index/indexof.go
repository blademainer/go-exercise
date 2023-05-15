package main

import (
	"strings"
)

func main() {
	extraParams := make(map[string]string)
	extraParam := "abc="
	idx := strings.Index(extraParam, "=")
	if idx < 0 {
		return
	}

	extraParams[extraParam[:idx]] = extraParam[idx+1:]
}
