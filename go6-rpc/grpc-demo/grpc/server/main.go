package main

import (
	"context"
	"log"
	"net"
	"server/protocol"

	"google.golang.org/grpc"
)

type Service struct {
	protocol.UnimplementedHelloServiceServer
}

func (s *Service) Hello(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	return &protocol.Response{
		Value: "Hello:" + req.Value,
	}, nil
}

func main() {
	// 首先是通过grpc.NewServer()构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务
	protocol.RegisterHelloServiceServer(grpcServer, new(Service))

	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	// 然后通过grpcServer.Serve(lis)在一个监听端口上提供gRPC服务
	grpcServer.Serve(lis)
}
