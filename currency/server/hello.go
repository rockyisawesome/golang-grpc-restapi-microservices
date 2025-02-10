package server

import (
	"context"
	protos "currency/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type HelloServer struct {
	loggs hclog.Logger
	protos.UnimplementedHelloWorldServer
}

func NewHelloServer(l hclog.Logger) *HelloServer {
	return &HelloServer{loggs: l}
}

func (hello *HelloServer) GetHello(_ context.Context, _ *protos.EmptyMessage) (*protos.HelloResponse, error) {
	hello.loggs.Info("I am Sending Hello World")
	return &protos.HelloResponse{Greet: "Hello World!!!-----!!!"}, nil
}
