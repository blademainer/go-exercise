package main

import (
	"crypto/tls"
	"fmt"
	tls2 "github.com/blademainer/go-exercise/demos/tls/tcp/tls"
	"io/ioutil"
)

func main() {
	config, err := tls2.NewTLSConfig(
		"demos/tls/key/ca.crt", "demos/tls/key/server.crt", "demos/tls/key/server.key",
	)
	if err != nil {
		panic(err.Error())
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:8080", config)
	if err != nil {
		panic(err.Error())
	}

	_, err = conn.Write([]byte("hello tls server"))
	if err != nil {
		panic(err.Error())
	}
	err = conn.CloseWrite()
	if err != nil {
		panic(err.Error())
	}

	all, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("reply: %v\n", string(all))
}
