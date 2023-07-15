package grpc

import (
	"context"

	"github.com/yuttasakcom/go-grpc-simple/protogen"
)

func (a *GrpcAdapter) StartPing(ctx context.Context, ping *protogen.Ping) (*protogen.Pong, error) {
	return a.pingpongService.StartPing(ctx, ping)
}
