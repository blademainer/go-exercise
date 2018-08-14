package main

import (
	"fmt"
	"sync"
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

	go func() {
		fmt.Println("Starting...")
		for {
			select {
			case s := <-dataCh:
				fmt.Printf("Receive i：%v \n", s)
				//doneCh <- true
			//default:
			//	fmt.Println("No data1...")
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
				//doneCh <- true
			//default:
			//	fmt.Println("No data1...")
			case <-closeCh:
				fmt.Printf("Closing...")
				wg.Done()
				return
			}
		}
	}()
	//go func() {
	//	fmt.Println("Starting...")
	//	for {
	//		select {
	//		case s := <-dataCh:
	//			fmt.Printf("Receive i：%v \n", s)
	//			doneCh <- true
	//		default:
	//			fmt.Println("No data2...")
	//			//case <-closeCh:
	//			//	fmt.Printf( "Closing...")
	//			//	//wg.Done()
	//			//	return
	//		}
	//	}
	//}()

	//for i := 0; i < 10000; i++ {
	//    <-doneCh
	//}
	//go func() {
	//	fmt.Println("Starting...")
	//	for {
	//		select {
	//		case s := <-dataCh:
	//			fmt.Printf("Receive i：%v \n", s)
	//		case <-closeCh:
	//			fmt.Printf( "Closing...")
	//			wg.Done()
	//			return
	//		}
	//	}
	//}()

	wg.Wait()
	close(dataCh)
	close(closeCh)

}
