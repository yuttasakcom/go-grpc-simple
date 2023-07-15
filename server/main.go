package main

import (
	"github.com/yuttasakcom/go-grpc-simple/internal/adapter/grpc"
	"github.com/yuttasakcom/go-grpc-simple/internal/application"
)

func main() {
	hs := &application.PingPongServerImpl{}
	grpcAdapter := grpc.NewGrpcAdapter(hs, 8080)
	grpcAdapter.Run()
}
