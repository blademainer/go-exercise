package main

import (
	"fmt"
	"runtime"
	"time"
)

//main.PrintStack()
//	/Volumes/ssddata/workspace/go/src/github.com/blademainer/go-exercise/demos/panic/reconver.go:11 +0x5b
//main.process.func1()
//	/Volumes/ssddata/workspace/go/src/github.com/blademainer/go-exercise/demos/panic/reconver.go:19 +0x7b
//panic(0x10abca0, 0x1160520)
//	/usr/local/go/src/runtime/panic.go:513 +0x1b9
//main.(*A).ttt(0x0)
//	/Volumes/ssddata/workspace/go/src/github.com/blademainer/go-exercise/demos/panic/reconver.go:30 +0x4b
//main.badFunc()
//	/Volumes/ssddata/workspace/go/src/github.com/blademainer/go-exercise/demos/panic/reconver.go:35 +0x2a
//main.process()
//	/Volumes/ssddata/workspace/go/src/github.com/blademainer/go-exercise/demos/panic/reconver.go:22 +0x3e
//main.main()
//	/Volumes/ssddata/workspace/go/src/github.com/blademainer/go-exercise/demos/panic/reconver.go:39 +0x22
func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

func process() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
			PrintStack()
		}
	}()
	badFunc()
}

type A struct {
	s string
}

func (a *A) ttt() {
	fmt.Println("s===", a.s)
}

func badFunc() {
	var a *A
	a.ttt()
}

func main() {
	process()
	time.Sleep(time.Second)
}
