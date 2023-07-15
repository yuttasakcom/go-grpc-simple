package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/yuttasakcom/go-grpc-simple/internal/port"
	"github.com/yuttasakcom/go-grpc-simple/protogen"
)

type GrpcAdapter struct {
	pingpongService port.PingPongServicePort
	grpcPort        int
	server          *grpc.Server
	protogen.PingPongServer
}

func NewGrpcAdapter(pingpongService port.PingPongServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		pingpongService: pingpongService,
		grpcPort:        grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v\n", a.grpcPort, err)
	}

	log.Printf("Starting gRPC server on port %d\n", a.grpcPort)

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	protogen.RegisterPingPongServer(grpcServer, a)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d: %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.GracefulStop()
}
