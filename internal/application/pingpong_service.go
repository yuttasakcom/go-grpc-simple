package application

import (
	"context"
	"fmt"

	pb "github.com/yuttasakcom/go-grpc-simple/protogen"
)

type PingPongServerImpl struct {
}

func (s *PingPongServerImpl) StartPing(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	fmt.Println("Ping Received")

	resp := pb.Pong{
		Id:      ping.Id,
		Message: ping.Message + " PONG",
	}

	return &resp, nil
}
