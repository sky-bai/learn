package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "helloworld/proto"
)

const (
	address     = "localhost:8090"
	defaultName = "world"
)

func main() {
	ctx1, _ := context.WithTimeout(context.Background(), 5*time.Second)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// Set up a connection to the server.
	conn, err := grpc.DialContext(ctx1, address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ！metadata 用于在服务间传递一些参数，注意后面它在交互中出现的位置
	//ctx = metadata.AppendToOutgoingContext(ctx, "metadata", "is metadata")

	md := metadata.Pairs(
		"key1", "val1",
		"key1", "val1-2", // "key1" will have map value []string{"val1", "val1-2"}
		"key2", "val2",
	)
	// 新建一个有 metadata 的 context
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Age: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

// http2.0 可以一个连接并发处理多个请求
// gRpc在三次握手之后，客户端/服务端会发送连接前言(Magic+SETTINGS)以确立协议和配置
// gRpc在传输数据过程中会设计滑动窗口(WINDOW_UPDATE)等流控策略
// gRpc附加信息基于HEADERS帧进行传递，具体的请求/响应数据存储在DATA帧中
// gRpc请求/响应结果分为HTTP和gRpc状态响应(grpc-status、grpc-message)两种类型
