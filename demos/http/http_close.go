package main

import (
	"fmt"
	"github.com/weaveworks/procspy"
	"net/http"
	"os"
	"time"
)

func main() {
	getpid := os.Getpid()
	fmt.Println("Pid: ", getpid)

	for i := 0; i < 1000; i++ {
		resp, err := http.DefaultClient.Get("https://baidu.com")
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		//if bytes, e := ioutil.ReadAll(resp.Body); e != nil {
		//	fmt.Println("Error: ", e.Error())
		//} else {
		//	fmt.Println(len(bytes))
		//}
		resp.Body.Close()
		time.Sleep(time.Second)
	}

	//connections()
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
