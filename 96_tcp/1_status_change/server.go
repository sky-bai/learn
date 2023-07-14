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

type XtConnProperty struct {
	Imei string `json:"imei"`
}

// DoConnectionBegin 创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {

	// 1.设置链接属性
	var x XtConnProperty
	conn.SetProperty("XtConnProperty", x)

	// 往链接里面发送消息
	//err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	//if err != nil {
	//	fmt.Println(err)
	//}
}

// DoConnectionLost 连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {

	// 业务1 清楚该链接对应设备的redis缓存
	property, err := conn.GetProperty("XtConnProperty")
	if err != nil {
		logger.Errorf(context.Background(), "GetProperty XtConnProperty error: %v", err)
	}
	xt, ok := property.(XtConnProperty)
	if !ok {
		logger.Errorf(context.Background(), "xt type error: %v", err)
	}

	// 链接断开的时候，删除redis中的设备缓存
	goredis.Backend.Del(context.Background(), "device:"+xt.Imei)
	goredis.Backend.ZRem(context.Background(), "online", xt.Imei)

	noAlive, err := conn.GetProperty("noAlive")
	if err != nil {
		logger.Errorf(context.Background(), "GetProperty noAlive error: %v", err)
	}

	msg := "正常退出"
	// 如果是异常退出
	if noAlive.(bool) {
		msg = "心跳超时了"
	}

	// 记录tcp日志
	_, err = model.LogTcpLogin.UpdateOne(context.Background(), bson.M{"conn_id": conn.GetConnUniqueId()}, bson.M{"$set": bson.M{"logout_time": time.Now().UTC(), "msg": msg}})
	if err != nil {
		logger.Error(context.Background(), "LogTcpLogin.InsertOne error:", err)
	}

}

// 用户自定义的远程连接不存活时的处理方法
func myOnRemoteNotAlive(conn ziface.IConnection) {

	//关闭链接
	conn.Stop()

	// 这里是异常退出
	conn.SetProperty("noAlive", true)

}

func main() {

	// 1.创建一个server句柄
	s := znet.NewServer("tcp")

	// 2.注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)

	// 3.注册链接断开时的hook函数
	s.SetOnConnStop(DoConnectionLost)

	// 4.配置路由
	s.AddRouter(znet.MsgType["XT"], &tcpHander.XtRouter{}, znet.TimeoutMiddleware(3))
	s.AddRouter(znet.MsgType["GPS"], &tcpHander.GpsRouter{})
	s.AddRouter(znet.MsgType["STATUS"], &tcpHander.StatusRouter{})
	s.AddRouter(znet.MsgType["UP"], &tcpHander.UpRouter{})
	s.AddRouter(znet.MsgType["ACCOFFMODEL"], &tcpHander.AccOffRouter{})
	s.AddRouter(znet.MsgType["ERROR"], &tcpHander.ErrorRouter{})
	s.AddRouter(znet.MsgType["CLOUDSTATE"], &tcpHander.CloudStateRouter{})

	// 5.启动心跳检测
	s.StartHeartBeatWithOption(1*time.Second, &ziface.HeartBeatOption{
		OnRemoteNotAlive: myOnRemoteNotAlive,
	})

	// 6.开启服务
	s.Serve()
}
