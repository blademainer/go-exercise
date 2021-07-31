package threadpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestThreadPool_Go(t *testing.T) {
	pool := NewThreadPool(10)
	count := 0
	size := 20
	hasErr := false
	lock := sync.Mutex{}
	for i := 0; i < size; i++ {
		err := pool.Go(
			func() {
				time.Sleep(1 * time.Second)
				lock.Lock()
				defer lock.Unlock()
				count++
				count++
			},
		)
		if err != nil {
			fmt.Println(err.Error())
			hasErr = true
			break
		}
	}
	if !hasErr {
		t.FailNow()
	}
}

func TestThreadPool_Run(t *testing.T) {
	pool := NewThreadPool(10)
	count := 0
	size := 2000
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		pool.Run(
			func() {
				defer wg.Done()
				lock.Lock()
				defer lock.Unlock()
				count++
			},
		)
	}
	wg.Wait()
	if count != size {
		t.Fatalf("count: %v not eq: %v", count, size)
	}
}
