package main

import (
	"log"
	"sync"
)

func main() {
	// disable GC so we can control when it happens.
	// defer debug.SetGCPercent(debug.SetGCPercent(-1))
	i := 0
	p := sync.Pool{
		New: func() interface{} {
			log.Printf("new: %v\n", i)
			i++
			return i
		},
	}
	if v := p.Get(); v != 1 {
		log.Fatalf("got %v; want 1", v)
	}
	if v := p.Get(); v != 2 {
		log.Fatalf("got %v; want 2", v)
	}

	// Make sure that the goroutine doesn't migrate to another P
	// between Put and Get calls.
	// debug.PrintStack()
	p.Put(42)
	if v := p.Get(); v != 42 {
		log.Fatalf("got %v; want 42", v)
	}
	// debug.PrintStack()


	if v := p.Get(); v != 3 {
		log.Fatalf("got %v; want 3", v)
	}
}
