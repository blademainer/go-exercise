package main

import (
	"context"
	"flag"
	"fmt"
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
	grpcPort   = flag.Int("grpcPort", 9200, "grpc port")
	healthPort = flag.Int("healthPort", 6666, "grpc health port")
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	logger.Infof("Starting app, version: %v", version)

	// shutdown functions
	shutdownFunctions := make([]func(context.Context), 0)

	ctx, cancel := context.WithCancel(context.Background())
	shutdownFunctions = append(
		shutdownFunctions, func(ctx context.Context) {
			cancel()
		},
	)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(
		func() error {
			// profiles := pprof.Profiles()

			httpServer := &http.Server{
				Addr:         fmt.Sprintf(":%d", *debugPort),
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
				Handler:      nil,
			}
			shutdownFunctions = append(
				shutdownFunctions, func(ctx context.Context) {
					err := httpServer.Shutdown(ctx)
					if err != nil {
						logger.Errorf("failed to shutdown pprof server! error: %v", err.Error())
					}
				},
			)

			logger.Infof("pprof server serving at :%d", *debugPort)

			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Errorf("failed to listen: %v", err.Error())
				return err
			}
			return nil
		},
	)

	// web server metrics
	g.Go(
		func() error {
			httpServer := &http.Server{
				Addr:         fmt.Sprintf(":%d", *httpPort),
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			}
			shutdownFunctions = append(
				shutdownFunctions, func(ctx context.Context) {
					err := httpServer.Shutdown(ctx)
					if err != nil {
						logger.Errorf("failed to shutdown pprof server! error: %v", err.Error())
					}
				},
			)
			logger.Infof("HTTP Metrics server serving at :%d", *httpPort)

			if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}

			return nil
		},
	)

	// gRPC Health Server
	healthServer := health.NewServer()
	g.Go(
		func() error {
			grpcHealthServer := grpc.NewServer()

			shutdownFunctions = append(
				shutdownFunctions, func(ctx context.Context) {
					healthServer.SetServingStatus(
						fmt.Sprintf("grpc.health.v1.%s", serviceName), healthpb.HealthCheckResponse_NOT_SERVING,
					)
					grpcHealthServer.GracefulStop()
				},
			)

			healthpb.RegisterHealthServer(grpcHealthServer, healthServer)

			haddr := fmt.Sprintf(":%d", *healthPort)
			hln, err := net.Listen("tcp", haddr)
			if err != nil {
				logger.Errorf("gRPC Health server: failed to listen, error: %v", err)
				os.Exit(2)
			}
			logger.Infof("gRPC health server serving at %s", haddr)
			return grpcHealthServer.Serve(hln)
		},
	)

	// gRPC server
	g.Go(
		func() error {
			addr := fmt.Sprintf(":%d", *grpcPort)
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				logger.Errorf("gRPC server: failed to listen, error: %v", err)
				os.Exit(2)
			}

			server := &server{}
			grpcServer := grpc.NewServer(
				// MaxConnectionAge is just to avoid long connection, to facilitate load balancing
				// MaxConnectionAgeGrace will torn them, default to infinity
				grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
			)
			pb.RegisterGreeterServer(grpcServer, server)
			shutdownFunctions = append(
				shutdownFunctions, func(ctx context.Context) {
					healthServer.SetServingStatus(
						fmt.Sprintf("grpc.health.v1.%s", serviceName), healthpb.HealthCheckResponse_NOT_SERVING,
					)
					grpcServer.GracefulStop()
				},
			)

			logger.Infof("gRPC server serving at %s", addr)

			healthServer.SetServingStatus(
				fmt.Sprintf("grpc.health.v1.%s", serviceName), healthpb.HealthCheckResponse_SERVING,
			)

			return grpcServer.Serve(ln)
		},
	)

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	logger.Warnf("received shutdown signal")

	// 创建一个新的Context，等待各个服务释放资源
	timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	for _, shutdown := range shutdownFunctions {
		shutdown(timeout)
	}

	err := g.Wait()
	if err != nil {
		logger.Errorf("server returning an error, error: %v", err)
		os.Exit(2)
	}
}
