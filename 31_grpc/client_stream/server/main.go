package main

import (
	pb "client_stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

// 2.我们需要实现grpc提供的接口
// 整体逻辑就是 先定义接口 ，然后实现接口，然后注册到grpc中，然后启动grpc服务

// 1.定义proto接口文件
// 2.实现grpc提供的接口
// 3.注册到grpc中

// SimpleService 定义我们的服务
type SimpleService struct {
	pb.UnimplementedStreamClientServer
}

// RouteList 实现RouteList方法
func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamData)
		if res.StreamData == "over" {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})

			// 如果是结束消息，就不再接收消息

		}
	}
}

// 第一个是如何判断流是否结束，这里我们使用了io.EOF，这是一个标准的错误类型，当流结束时，会返回这个错误。
// 熄火/断线在redis中将它的地理位置删除 是判断redis是否有值，来判断是否在线 那是否可以在设备端来发送我是否在线的消息。

const (
	// Address 监听地址
	Address string = ":9900"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	// todo 始终有一个server实例，然后将服务注册到server中，然后启动 2.这里可以把server注册到etcd中

	// 1.新建gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 2.在gRPC服务器注册我们的服务 也可以注册到etcd上面
	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})

	// 3.用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
