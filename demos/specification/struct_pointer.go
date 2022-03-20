package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type P struct {
	Name  string
	Age   uint8
	cache *fieldCache
}

type fieldCache struct {
	value atomic.Value // map[reflect.Type][]field
	mu    sync.Mutex   // used only by writers
}

func (p *P) Process() {
	fmt.Printf("Process &p: %p \n", p)
	printAndReadFields(*p)
}

func printAndReadFields(p P) {
	fmt.Printf("printAndReadFields &p: %p \n", &p)

	if p.cache == nil {
		fmt.Println("init cache...")
		p.cache = &fieldCache{}
	}
	// update cache...
}

func main() {
	p := &P{}
	p.Name = "张三"
	p.Age = 18
	p.Process()
	p.Process()
	p.Process()
}
