package main

import (
	"net"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		panic("args is nil!")
	}
	file, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}

	listener, e := net.Listen("tcp", "127.0.0.1:8080")
	if e != nil {
		panic(e)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
	e = http.Serve(listener, serve(*file))
	if e != nil {
		panic(e)
	}
}

func serve(file os.File) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, file.Name())
	}
}
