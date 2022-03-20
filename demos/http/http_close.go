package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/weaveworks/procspy"
)

func main() {
	getpid := os.Getpid()
	fmt.Println("Pid: ", getpid)

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := http.DefaultClient.Get("http://192.168.25.140")
				if err != nil {
					fmt.Println("Error: ", err.Error())
					return
				}
				// if bytes, e := ioutil.ReadAll(resp.Body); e != nil {
				//	fmt.Println("Error: ", e.Error())
				// } else {
				//	fmt.Println(len(bytes))
				// }
				resp.Body.Close()
				time.Sleep(100 * time.Second)
			}()
		}
	}

	wg.Wait()

	// connections()
}

func connections() {
	lookupProcesses := true
	cs, err := procspy.Connections(lookupProcesses)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("TCP Connections:\n")
	for c := cs.Next(); c != nil; c = cs.Next() {
		fmt.Printf(" - %v\n", c)
	}
}
