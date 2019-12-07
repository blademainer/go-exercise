package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestDo(t *testing.T) {
	a := make(map[int]int)
	rw := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			sum := i + j
			k := sum % 10
			go func() {
				exists := -1
				Do(rw,
					func() bool {
						var found bool
						exists, found = a[k]
						return !found
					},
					func() {
						fmt.Printf("found: %v replace with: %v\n", exists, sum)
						a[k] = sum
					},
				)
				wg.Done()
			}()

		}
	}
	wg.Wait()

}

func TestDo2(t *testing.T) {
	a := make(map[int]int)
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				k := (i + j) % 10
				v := i + j
				exists := -1

				Do(rw,
					func() bool {
						var found bool
						exists, found = a[k]
						return !found
					},
					func() {
						fmt.Printf("found: %v replace with: %v\n", exists, v)
						a[k] = v
					},
				)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func TestDo3(t *testing.T) {
	a := make(map[int]int)
	wg := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(100)
		go func() {
			for j := 0; j < 100; j++ {
				k := (i + j) % 10
				v := i + j
				f := func() {
					rwMutex.RLock()
					defer rwMutex.RUnlock()
					if _, found := a[k]; !found {
						//       code block start								       code block end
						//             ⬇													  ⬇
						// ReadLock -> {  ReadUnlock -> Lock -> write() -> UnLock -> ReadLock } -> ReadUnlock

						rwMutex.RUnlock()
						defer rwMutex.RLock()
						rwMutex.Lock()
						defer rwMutex.Unlock()
						fmt.Println("write: ", k, " v: ", v)

						a[k] = v
					}
				}
				f()
				wg.Done()
			}
		}()
	}
	wg.Wait()
}

func TestDo4(t *testing.T) {
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	sum := 0

	for i := 0; i < 1000; i++ {
		wg.Add(100)
		go func() {
			for j := 0; j < 100; j++ {
				// bad
				Do(rw,
					func() bool {
						fmt.Println(sum)
						return true
					},
					func() {
						sum++
					},
				)
				wg.Done()
			}
		}()
	}
	wg.Wait()
	fmt.Println("result: ", sum)
}

func TestDo5(t *testing.T) {
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	sum := 0

	for i := 0; i < 1000; i++ {
		wg.Add(100)
		go func() {
			for j := 0; j < 100; j++ {
				//rw.Lock()
				DoFunc(rw, func() {
					sum++
				})
				//rw.Unlock()
				wg.Done()
			}
		}()
	}
	wg.Wait()
	fmt.Println("result: ", sum)
}

func TestDo6(t *testing.T) {
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	sum := 0

	for i := 0; i < 1000; i++ {
		wg.Add(100)
		go func() {
			for j := 0; j < 100; j++ {
				rw.Lock()
				f := func() {
					sum++
				}
				f()
				rw.Unlock()
				wg.Done()
			}
		}()
	}
	wg.Wait()
	fmt.Println("result: ", sum)
}

func DoFunc(rw sync.RWMutex, f func()) {
	rw.Lock()
	f()
	rw.Unlock()
}
