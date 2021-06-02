package main

import (
	"crypto/tls"
	"fmt"
	tls2 "github.com/blademainer/go-exercise/demos/tls/tcp/tls"
	"io/ioutil"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}
	config, err := tls2.NewTLSConfig(
		"demos/tls/key/ca.crt", "demos/tls/key/server.crt", "demos/tls/key/server.key",
	)
	if err != nil {
		panic(err.Error())
	}
	listener := tls.NewListener(listen, config)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	all, err2 := ioutil.ReadAll(conn)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Printf("get request: %v\n", string(all))
	data := []byte("ok")
	_, err2 = conn.Write(data)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Printf("reply: %v\n", string(data))
	err2 = conn.Close()
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
}
