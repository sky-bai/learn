package main

import (
	pb "client_stream/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

// routeList 调用服务端RouteList方法
func routeList() {
	//调用服务端RouteList方法，获流
	stream, err := streamClient.RouteList(context.Background()) // 1.拿到客户端的流
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	for n := 0; n < 5; n++ { // 使用定时器去发
		//向流中发送消息
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}

// Address 连接地址
const Address string = ":8000"

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
