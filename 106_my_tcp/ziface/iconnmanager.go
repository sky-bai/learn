package ziface

/*
 链接管理模块 增加链接个数的限定，如果超过一定量的客户端个数，Zinx为了保证后端的及时响应，而拒绝链接请求。
 ConnManager中，其中用一个map来承载全部的连接信息，key是连接ID，value则是连接本身。其中有一个读写锁connLock主要是针对map做多任务修改时的保护作用。
*/

type IConnManager interface {
	Add(conn IConnection)                   //添加链接
	Remove(conn IConnection)                //删除连接
	Get(connID string) (IConnection, error) //利用ConnID获取链接
	Len() int                               //获取当前连接
	ClearConn()                             //删除并停止所有链接
}
