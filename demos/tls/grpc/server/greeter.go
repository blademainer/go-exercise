package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	testdata "github.com/blademainer/go-exercise/demos/tls/key"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		main1()
	}()
	// wg.Add(1)
	// go func() {
	//	defer wg.Done()
	//	main2() // bad
	// }()
	wg.Wait()
}

func main1() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile(testdata.Path("server.crt"), testdata.Path("server.key"))
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	gs := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(gs, &server{})
	panic(gs.Serve(lis))
}

func main2() {
	cert, err := ioutil.ReadFile(testdata.Path("server.crt"))
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	key, err := ioutil.ReadFile(testdata.Path("server.key"))
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	certificate, err := tls.LoadX509KeyPair(string(cert), string(key))

	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile(testdata.Path("ca.crt"))
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	if ok := pool.AppendCertsFromPEM(caCrt); !ok {
		fmt.Println("failed to append")
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "", 50052))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	creds := credentials.NewTLS(
		&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
			ClientCAs:    pool,
		},
	)
	gs := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(gs, &server{})
	panic(gs.Serve(lis))
}
