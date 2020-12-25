/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Binary client is an example client.
package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

var p = sync.Pool{
	New: func() interface{} {
		conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		return conn
	},
}

func test() {
	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	get := p.Get()
	defer p.Put(get)
	c := pb.NewGreeterClient(get.(*grpc.ClientConn))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Printf("error: %v", err.Error())
	} else {
		log.Printf("Greeting: %s", r.Message)
	}
}

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				test()
			}
		}()
	}
	wg.Wait()
}
