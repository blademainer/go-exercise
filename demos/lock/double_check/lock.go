package main

import (
	"fmt"
	"sync"
)

type store struct {
	sync.RWMutex
	kv map[int]int
}

// maybe panic
func (store *store) getOrInitWithoutRLock(k int, v int) int {
	exists, found := store.kv[k]
	if !found {
		store.Lock()
		defer store.Unlock()
		if exists, found := store.kv[k]; found {
			fmt.Printf("key: %v is already inited by another thread!\n", k)
			return exists
		}
		store.kv[k] = v
		fmt.Printf("key: %v is inited by this thread!\n", k)
	}
	return exists
}

func (store *store) getOrInit(k int, v int) int {
	store.RLock()
	defer store.RUnlock()
	exists, found := store.kv[k]
	if !found {
		store.RUnlock()
		defer store.RLock()
		store.Lock()
		defer store.Unlock()
		if exists, found := store.kv[k]; found {
			fmt.Printf("key: %v is already inited by another thread!\n", k)
			return exists
		}
		store.kv[k] = v
		fmt.Printf("key: %v is inited by this thread!\n", k)
	}
	return exists
}

func main() {
	s := &store{kv: make(map[int]int)}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				k := (i + j) % 10
				v := k
				s.getOrInit(k, v)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	s2 := &store{kv: make(map[int]int)}
	wg2 := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg2.Add(100)
		go func() {
			for j := 0; j < 100; j++ {
				k := (i + j) % 10
				v := k
				s2.getOrInit(k, v)
				wg2.Done()
			}
		}()
	}
	wg2.Wait()
}
