package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync/atomic"
	"time"
)

var index int32 = 0


func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:16060", nil))
	}()

	s := &http.Server{
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
		Addr:":8080",
	}


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ii := atomic.AddInt32(&index, 1)
		fmt.Printf("[%v]HandleApi Request@%v \n", ii, time.Now())

		closer := request.Body
		if closer != nil {
			//defer closer.Close()
		}
		bytes, e := ioutil.ReadAll(closer)
		if e != nil {
			fmt.Println(e.Error())
			_, e := writer.Write([]byte("ERROR"))
			if e != nil {
				fmt.Println(e.Error())
			}
			return
		}
		fmt.Printf("Body: %v \n", string(bytes))

		_, e = writer.Write([]byte("hello"))
		if e != nil {
			fmt.Println(e.Error())
		}
		fmt.Printf("[%v]HandleApi Response@%v \n", ii, time.Now())
	})




	log.Println(s.ListenAndServe())
}
