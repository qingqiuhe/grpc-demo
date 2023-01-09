package main

import (
	"context"
	pb "github.com/qingqiuhe/grpc-gateway-demo/helloworld"
)

const port = ":50051"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.)