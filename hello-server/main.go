package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-study/hello-server/proto"
	"net"
)

// hello Server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic("port open failed")
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		panic("service open failed")
	}

}
