package main

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	ms := &jsonpb.Marshaler{

	}
	r := &helloworld.HelloRequest{Name: "zhagnsan&asd>=123"}
	str, err := ms.MarshalToString(r)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(str)
}
