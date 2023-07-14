package main

import (
	"backend/bin/tcp/tcpHander"
	"backend/bin/tcp/ziface"
	"backend/bin/tcp/znet"
	"backend/lib/goredis"
	"backend/lib/logger"
	"backend/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// DoConnectionLost 连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {

	// 1.链接断开的时候，删除redis中的设备缓存
	goredis.Backend.Del(context.Background(), "device:"+conn.GetImei())
	goredis.Backend.ZRem(context.Background(), "online", conn.GetImei())

	msg := ""

	noAlive, exist := conn.Get("noAlive")
	if !exist {
		msg = "正常退出"
	} else {
		msg = noAlive.(string)
	}

	// 记录tcp日志
	_, err := model.LogTcpLogin.UpdateOne(context.Background(), bson.M{"conn_id": conn.GetConnID()}, bson.M{"$set": bson.M{"logout_time": time.Now().UTC(), "msg": msg}})
	if err != nil {
		logger.Error(context.Background(), "LogTcpLogin.InsertOne error:", err)
	}

}

// 用户自定义的远程连接不存活时的处理方法
func myOnRemoteNotAlive(conn ziface.IConnection) {
	// 这里是异常退出
	conn.Set("noAlive", "心跳超时了")
}

func main() {

	// 1.创建一个server句柄
	s := znet.NewServer()

	// 2.serve层注册链接断开时的hook函数
	s.SetOnConnStop(DoConnectionLost)

	// 3.配置AddHandler
	s.AddHandler("XT", tcpHander.XtHandlerFunc)
	s.AddHandler("GPS", tcpHander.GpsHandlerFunc)
	s.AddHandler("STATUS", tcpHander.StatusHandlerFunc)
	s.AddHandler("UP", tcpHander.UpHandlerFunc)
	s.AddHandler("ACCOFFMODEL", tcpHander.AccOffHandlerFunc)
	s.AddHandler("ERROR", tcpHander.ErrorHandlerFunc)
	s.AddHandler("CLOUDSTATE", tcpHander.CloudStateHandlerFunc)

	// 4.启动心跳检测
	s.StartHeartBeatWithOption(1*time.Second, &ziface.HeartBeatOption{
		OnRemoteNotAlive: myOnRemoteNotAlive,
	})

	// 5.开启服务
	s.Serve()
}
