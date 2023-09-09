package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/natefinch/lumberjack.v2"
	pb "helloworld/proto"
	"log"
	"net"
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

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(ZapInterceptor()),
		)))
	defer s.Stop()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//  以json 输出 带有时间戳

// 在拦截器中 进行日志的打印配置
// todo 这里记录的是什么日志

func ZapInterceptor() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "log/debug.log",
		MaxSize:   1024, //MB
		LocalTime: true,
	})
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.NewAtomicLevel(),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}

// 那按照道理来说 日志都是服务端在打印 设备日志 c++日志 go日志 全部统一搜集
// http 日志 和 grpc 日志都搜集起来

// 如何替换为etcd

// 相当于之前的dns是将域名转换成ip地址，这里一种转换， 现在服务器是自己配置的了，转换规则由自己定义 etcd 保存
