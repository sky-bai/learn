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

// 将消息队列（MQ）替换为gRPC通信是一项较大的改动，可能需要进行全面的系统设计和开发。因此，如果您的现有系统已经能够满足您的需求，那么将MQ替换为gRPC可能会导致一定的风险和影响，需要您进行充分的评估和测试。
//
//如果您决定使用gRPC来替换MQ，应该考虑以下因素：
//
//数据大小：gRPC通常比MQ更适合小型消息的传输，而对于大型消息，MQ可能比gRPC更适合。
//
//可靠性：MQ可以提供更高的消息可靠性和持久性，因为消息在存储和传输过程中会进行持久化和复制，而gRPC则不具备这些特性。如果您需要更高的消息可靠性和持久性，可能需要考虑使用专门的MQ系统。
//
//实现复杂度：将MQ替换为gRPC可能需要更多的系统设计和开发工作，因为gRPC需要定义数据结构和服务接口，而MQ则只需要定义消息格式即可。
//
//性能：gRPC通常比MQ更快速和高效，但是具体性能取决于具体实现和网络环境等因素。
//
//因此，在考虑将MQ替换为gRPC时，应该充分评估您的需求、可用资源和技术能力等因素，以确保您做出了正确的决策。
