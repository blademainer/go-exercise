package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	ms := &jsonpb.Marshaler{}
	r := &helloworld.HelloRequest{Name: "zhagnsan&asd=123"}
	str, err := ms.MarshalToString(r)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("jsonpb: ", str)

	marshal, err := json.Marshal(r)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("std json: ", string(marshal))

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("disable escape json: ", buf.String())
}
