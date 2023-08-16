package main

import (
	pb "client_stream/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

// 情景模拟：客户端大量数据上传到服务端

// routeList 调用服务端RouteList方法
func routeList() {
	//调用服务端RouteList方法，获流
	stream, err := streamClient.RouteList(context.Background()) // 1.拿到客户端的流
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	// 使用定时器去发 //向流中发送消息
	for i := 0; i < 5; i++ {
		err = stream.Send(&pb.StreamRequest{StreamData: "stream client rpc "})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	err = stream.Send(&pb.StreamRequest{StreamData: "over"})
	if err != nil {
		log.Fatalf("stream request err: %v", err)
	}

	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}

// Address 连接地址
const Address string = ":9900"

var streamClient pb.StreamClientClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClientClient(conn)
	routeList()
}
