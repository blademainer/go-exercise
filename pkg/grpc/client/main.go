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
	"github.com/blademainer/go-exercise/pkg/grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

func main() {
	flag.Parse()

	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Set up a connection to the server.
	conn, err := grpc.DialContext(timeout, *addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			log.Printf("failed to close connection: %s", e)
		}
	}()
	client := proto.NewCurlServiceClient(conn)

	ctx, cancelRpc := context.WithTimeout(context.Background(), time.Second)
	defer cancelRpc()

	// may caused error: context deadline exceeded. because of the grpc context is timeout.
	r, err := client.Curl(ctx, &proto.CurlRequest{Url: "https://dldir1.qq.com/weixin/android/weixin7013android1640.apk"})
	if err != nil {
		log.Printf("error: %v", err.Error())
	} else {
		log.Printf("Response data len: %v", len(r.Data))
	}
	r, err = client.Curl(ctx, &proto.CurlRequest{}) // rpc error: code = InvalidArgument desc = invalid CurlRequest.Url: value must be absolute
	if err != nil {
		log.Printf("error: %v", err.Error())
	} else {
		log.Printf("Response data len: %v", len(r.Data))
	}
}
