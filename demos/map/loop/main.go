package main

import "fmt"

// go build -o rwop.bin -gcflags -S ./

func main() {
	ageMp := make(map[string]int)
	ageMp["qcrao"] = 18
	for name := range ageMp {
		ageMp[name+name] = 1
		delete(ageMp, name)
		fmt.Println(name)
	}
	fmt.Println(ageMp)
}
