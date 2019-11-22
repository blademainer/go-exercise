package main

import (
	"fmt"
	"runtime"
)

func A() {
	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs)
	for _, pc := range pcs[:i] {
		funcPC := runtime.FuncForPC(pc)
		println(funcPC.Name())
	}

	PrintStack()

}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

func main() {
	A()
}
