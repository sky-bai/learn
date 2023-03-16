package znet

import (
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// Start 启动服务器
func (s *Server) Start() {
	// TODO
	go func() {
		add, err := net.ResolveTCPAddr(s.IPVersion, s.IP+":"+string(s.Port))
		if err != nil {
			panic(err)
		}
		listener, err := net.ListenTCP(s.IPVersion, add)
		if err != nil {
			panic(err)
			return
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 获取到连接后，做一些业务处理
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop() {
	// TODO

}

func (s *Server) Serve() {
	// TODO
	s.Start()

}
