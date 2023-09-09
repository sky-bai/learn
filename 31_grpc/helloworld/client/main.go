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
	// 如何验证超时设置昵

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// 之前是ip加端口
	// Set up a connection to the server.
	conn, err := grpc.DialContext(ctx1, address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	// 接口定义的方法

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

// trailer 用于在请求/响应结束后传递一些额外的信息，比如错误信息

// 现在一元请求/响应模式已经足够应对大部分场景，但是在某些场景下，我们需要双向流式传输，比如聊天室、视频直播等场景，这时候就需要使用流式传输模式了
// 心跳包 应该是每个设备都要服务器发送xt 请求数据 当前设备号 设备的数据
// 客户端流式传输 设备端是c++ 服务端是go 使用grpc通信

// 一元请求/响应模式

// http log 和 grpc log 统一 统一是什么意思

// 服务之间使用链路追踪

// 1.测试超时
// 2.dial时参数的限制
