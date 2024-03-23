package server

import "github.com/go-kratos/kratos/v2/transport/grpc"

func NewGRPCServer() *grpc.Server {
	svr := grpc.NewServer()
	return svr
}
