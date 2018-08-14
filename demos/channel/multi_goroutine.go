package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	dataCh := make(chan int, 1024)
	closeCh := make(chan bool, 2)
	//doneCh := make(chan bool, 1024)
	go func() {
		for i := 0; i < 10000; i++ {
			dataCh <- i
		}
		closeCh <- true
		closeCh <- true
		//wg.Done()
	}()

	count := int32(0)
	go func() {
		fmt.Println("Starting...")
		for {
			select {
			case s := <-dataCh:
				fmt.Printf("Receive i：%v \n", s)
				atomic.AddInt32(&count, 1)
			case <-closeCh:
				fmt.Printf("Closing...")
				wg.Done()
				return
			}
		}
	}()
	go func() {
		fmt.Println("Starting...")
		for {
			select {
			case s := <-dataCh:
				fmt.Printf("Receive i：%v \n", s)
				atomic.AddInt32(&count, 1)
			case <-closeCh:
				fmt.Printf("Closing...")
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
	close(dataCh)
	close(closeCh)
	fmt.Println("count:", count)

}
