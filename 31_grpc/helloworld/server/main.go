package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "helloworld/proto"
)

const (
	port = ":8090"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	// 1.传递ctx参数
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("metadata:", md["key1"]) // 取出来是个切片数组 // metadata一般放什么数据 和 请求里面的参数有什么区别
	}
	return &pb.HelloReply{Message: "Hello111 ", Address: "123"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	//defer s.Stop()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
