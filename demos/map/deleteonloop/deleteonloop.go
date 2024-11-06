package main

import "fmt"

func main() {
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "",
		"d": "4",
		"e": "",
	}

	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}
	fmt.Println(m)
}
