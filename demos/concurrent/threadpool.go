package main

import (
	"fmt"
	"sync"
	"time"
)

func run(f func(), done chan bool) {
	f()
	done <- true
}

type ThreadPool struct {
	i    int
	Size int
	lock sync.Mutex
}

func server(threadSize uint8, fs chan func()) {
	i := 0
	done := make(chan bool)
	for {
		f := <-fs
		// fmt.Printf("current: %d >= max: %d \n", i, threadSize)
		fmt.Printf("index: %d fs=== %s \n", i, f)
		go run(f, done)
		i++
	}
}

func increase() {

}

func start() chan func() {
	ch := make(chan func())
	go server(10, ch)
	return ch
}

func main() {
	fc := start()
	a := func() {
		fmt.Println("Sleeping...")
		time.Sleep(time.Second * 1)
		fmt.Println("Done.")
	}
	for i := 0; i <= 100; i++ {
		fc <- a
	}

	time.Sleep(time.Second * 100)
}
