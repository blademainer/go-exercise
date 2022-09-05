package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = sync.RWMutex{}
var m = make(map[string]string)

func read() {
	lock.RLock() // may deadlock with: "fatal error: all goroutines are asleep - deadlock!"
	if v, ok := m["name"]; ok {
		fmt.Println("read value: ", v)
		// missing release lock
	} else {
		fmt.Println("not found name")
		lock.RUnlock()
	}
}

func write() {
	lock.Lock() // may deadlock with: "fatal error: all goroutines are asleep - deadlock!"
	defer lock.Unlock()
	fmt.Println("writing")
	m["name"] = "hello"
	fmt.Println("written ok")
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			read()
		}
	}()
	go func() {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Millisecond)
			write()
		}
	}()
	wg.Wait()
}
