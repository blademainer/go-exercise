package main

import "fmt"

// go build -o rwop.bin -gcflags -S ./

func main() {
	ageMp := make(map[string]int)
	ageMp["qcrao"] = 18
	for name := range ageMp {
		delete(ageMp, name)
		ageMp[name+name] = 1
		fmt.Println(name)
	}
	fmt.Println(ageMp)
}
