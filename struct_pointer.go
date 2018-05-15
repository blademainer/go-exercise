package main

import (
	"sync/atomic"
	"sync"
	"fmt"
	"reflect"
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



func (p *P) Process(){
	fmt.Printf("%p \n", p)
	printAndReadFields(*p, *p)
}

func printAndReadFields(p P, entity interface{}) {
	fmt.Printf("%p \n", &p)

	fields := (&p).cachedTypeFields(reflect.ValueOf(entity))
	fmt.Println("fields: ", fields)
}

type field struct {
	name string
	typ reflect.Value
}

func (p *P) cachedTypeFields(t reflect.Value) []field {
	fmt.Printf("%p \n", p)
	cache := p.cache.value.Load().(map[reflect.Type][]field)
	fields := cache[t.Type()]
	if fields != nil {
		return fields
	}
	readFields := readFields(t.Type())
	p.cache.mu.Lock()
	p.cache.value.Store(readFields)
	p.cache.mu.Unlock()
	return readFields
}

func readFields(t reflect.Type) []field {
	return nil
}

func main() {
	p := &P{}
	p.Name = "张三"
	p.Age = 18
	p.Process()
}
