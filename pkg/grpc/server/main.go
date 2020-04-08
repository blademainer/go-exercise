package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/blademainer/go-exercise/pkg/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blademainer/commons/pkg/logger"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

const serviceName = "myserviced"

var (
	version = "no version"

	debugPort  = flag.Int("debugPort", 16161, "debug port")
	httpPort   = flag.Int("httpPort", 8888, "http port")
	grpcPort   = flag.Int("grpcPort", 50052, "grpc port")
	healthPort = flag.Int("healthPort", 6666, "grpc health port")
)

type server struct {
	client http.Client
}

func (s *server) Curl(ctx context.Context, request *proto.CurlRequest) (*proto.Response, error) {
	err := request.Validate()
	if err != nil {
		logger.Errorf("failed to validate request: %v error: %v", request, err.Error())
		e := status.Error(codes.InvalidArgument, err.Error())
		return nil, e
	}
	// http request with grpc context
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, request.Url, nil)
	if err != nil {
		logger.Errorf("failed to create http request, url: %v error: %v", request.Url, err.Error())
		err := status.Error(codes.Internal, "failed to download file")
		return nil, err
	}
	resp, err := s.client.Do(r)
	if err != nil {
		logger.Errorf("failed to get file via http: %v", err.Error())
		err := status.Error(codes.Internal, "failed to download file")
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			err2 := resp.Body.Close()
			if err2 != nil {
				logger.Errorf("close error: %v", err2)
			}
		}
	}()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// may caused error: context deadline exceeded. because of the grpc context is timeout.
		logger.Errorf("failed to read response, error: %v", err.Error())
		e := status.Error(codes.Internal, err.Error())
		return nil, e
	}

	response := &proto.Response{}
	response.Data = all
	return response, nil
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	if in.Name == "world" {
		return nil, status.Error(codes.InvalidArgument, "forbidden world")
	}
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	logger.Infof("Starting app, version: %v", version)

	// shutdown functions
	shutdownFunctions := make([]func(), 0)

	ctx, cancel := context.WithCancel(context.Background())
	shutdownFunctions = append(shutdownFunctions, cancel)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		//profiles := pprof.Profiles()

		httpServer := &http.Server{
			Addr:         fmt.Sprintf(":%d", *debugPort),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      nil,
		}
		shutdownFunctions = append(shutdownFunctions, func() {
			err := httpServer.Shutdown(ctx)
			if err != nil {
				logger.Errorf("failed to shutdown pprof server! error: %v", err.Error())
			}
		})

		logger.Infof("pprof server serving at :%d", *debugPort)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("failed to listen: %v", err.Error())
			return err
		}
		return nil
	})

	// web server metrics
	g.Go(func() error {
		httpServer := &http.Server{
			Addr:         fmt.Sprintf(":%d", *httpPort),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		shutdownFunctions = append(shutdownFunctions, func() {
			err := httpServer.Shutdown(ctx)
			if err != nil {
				logger.Errorf("failed to shutdown pprof server! error: %v", err.Error())
			}
		})
		logger.Infof("HTTP Metrics server serving at :%d", *httpPort)

		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	// gRPC Health Server
	healthServer := health.NewServer()
	g.Go(func() error {
		grpcHealthServer := grpc.NewServer()

		shutdownFunctions = append(shutdownFunctions, func() {
			healthServer.SetServingStatus(fmt.Sprintf("grpc.health.v1.%s", serviceName), healthpb.HealthCheckResponse_NOT_SERVING)
			grpcHealthServer.GracefulStop()
		})

		healthpb.RegisterHealthServer(grpcHealthServer, healthServer)

		haddr := fmt.Sprintf(":%d", *healthPort)
		hln, err := net.Listen("tcp", haddr)
		if err != nil {
			logger.Errorf("gRPC Health server: failed to listen, error: %v", err)
			os.Exit(2)
		}
		logger.Infof("gRPC health server serving at %s", haddr)
		return grpcHealthServer.Serve(hln)
	})

	// gRPC server
	g.Go(func() error {
		addr := fmt.Sprintf(":%d", *grpcPort)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			logger.Errorf("gRPC server: failed to listen, error: %v", err)
			os.Exit(2)
		}

		server := &server{
			client: http.Client{
				Timeout: 1200 * time.Second,
			},
		}
		grpcServer := grpc.NewServer(
			// MaxConnectionAge is just to avoid long connection, to facilitate load balancing
			// MaxConnectionAgeGrace will torn them, default to infinity
			grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
		)
		pb.RegisterGreeterServer(grpcServer, server)
		proto.RegisterCurlServiceServer(grpcServer, server)
		shutdownFunctions = append(shutdownFunctions, func() {
			healthServer.SetServingStatus(fmt.Sprintf("grpc.health.v1.%s", serviceName), healthpb.HealthCheckResponse_NOT_SERVING)
			grpcServer.GracefulStop()
		})

		logger.Infof("gRPC server serving at %s", addr)

		healthServer.SetServingStatus(fmt.Sprintf("grpc.health.v1.%s", serviceName), healthpb.HealthCheckResponse_SERVING)

		return grpcServer.Serve(ln)
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	logger.Warnf("received shutdown signal")

	for _, shutdown := range shutdownFunctions {
		shutdown()
	}

	err := g.Wait()
	if err != nil {
		logger.Errorf("server returning an error, error: %v", err)
		os.Exit(2)
	}
}
