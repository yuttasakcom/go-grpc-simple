package port

import (
	"context"

	pb "github.com/yuttasakcom/go-grpc-simple/protogen"
)

type PingPongServicePort interface {
	StartPing(ctx context.Context, ping *pb.Ping) (*pb.Pong, error)
}
