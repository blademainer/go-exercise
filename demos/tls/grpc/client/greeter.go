package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	testdata "github.com/blademainer/go-exercise/demos/tls/key"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	address2     = "localhost:50052"
	defaultName = "world"
)

func main() {
	main1()
	//main2()
}

func main1() {
	creds, err := credentials.NewClientTLSFromFile(testdata.Path("ca.crt"), "localhost")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}


	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func main2() {
	// Create tls based credential.
	cert, err := tls.LoadX509KeyPair(testdata.Path("client.crt"), testdata.Path("client.key"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	caCrt, err := ioutil.ReadFile(testdata.Path("ca.crt"))
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCrt) {
		err := errors.New("fail to append ca")
		fmt.Println(err.Error())
		return
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "localhost",
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})

	// Set up a connection to the server.
	conn, err := grpc.Dial(address2, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
