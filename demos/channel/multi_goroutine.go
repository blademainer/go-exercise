package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	dataCh := make(chan int, 1000)
	closeCh := make(chan bool, 2)
	//doneCh := make(chan bool, 1024)
	go func() {
		for i := 0; i < 10000; i++ {
			dataCh <- i
		}
		fmt.Println("Done...")
		closeCh <- true
		closeCh <- true
		//wg.Done()
		//close(dataCh)
	}()

	count := int32(0)
	go func() {
		defer wg.Done()
		fmt.Println("Starting...")
		for {
			select {
			case s := <-dataCh:
				fmt.Printf("Receive i：%v \n", s)
				atomic.AddInt32(&count, 1)
			case <-closeCh:
				fmt.Printf("Closing... cur length: %d \n", len(dataCh))
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		defer func() {
			for len(dataCh) > 0{
				fmt.Printf("Read least data: %d \n", <-dataCh)
				atomic.AddInt32(&count, 1)
			}
		}()
		fmt.Println("Starting...")
		for {
			select {
			case s := <-dataCh:
				fmt.Printf("Receive i：%v \n", s)
				atomic.AddInt32(&count, 1)
			case <-closeCh:
				fmt.Printf("Closing... cur length: %d \n", len(dataCh))
				return
			}
		}
	}()

	wg.Wait()
	close(dataCh)
	close(closeCh)
	fmt.Println("count:", count)

}
