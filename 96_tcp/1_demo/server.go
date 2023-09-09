package main

import (
	// "golang.org/x/sys/unix"
	"log"
	"net"
	"net/http"
	"reflect"
	"sync"
)

var epoller *epoll

func main() {
	//setLimit()
	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()
	//epoller, err = MkEpoll()
	//if err != nil {
	//	panic(err)
	//}
	//go start()
	//for {
	//	conn, e := ln.Accept()
	//	if e != nil {
	//		if ne, ok := e.(net.Error); ok && ne.Temporary() {
	//			log.Printf("accept temp err: %v", ne)
	//			continue
	//		}
	//		log.Printf("accept err: %v", e)
	//		return
	//	}
	//	if err := epoller.Add(conn); err != nil {
	//		log.Printf("failed to add connection %v", err)
	//		conn.Close()
	//	}
	//}
}

//func start() {
//	var buf = make([]byte, 8)
//	for {
//		connections, err := epoller.Wait()
//		if err != nil {
//			log.Printf("failed to epoll wait %v", err)
//			continue
//		}
//		for _, conn := range connections {
//			if conn == nil {
//				break
//			}
//			if _, err := conn.Read(buf); err != nil {
//				if err := epoller.Remove(conn); err != nil {
//					log.Printf("failed to remove %v", err)
//				}
//				conn.Close()
//			}
//		}
//	}
//}

type epoll struct {
	fd          int
	connections map[int]net.Conn
	lock        *sync.RWMutex
}

//	func MkEpoll() (*epoll, error) {
//		fd, err := unix.EpollCreate1(0)
//		if err != nil {
//			return nil, err
//		}
//		return &epoll{
//			fd:          fd,
//			lock:        &sync.RWMutex{},
//			connections: make(map[int]net.Conn),
//		}, nil
//	}
//
//	func (e *epoll) Add(conn net.Conn) error {
//		// Extract file descriptor associated with the connection
//		fd := socketFD(conn)
//		err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
//		if err != nil {
//			return err
//		}
//		e.lock.Lock()
//		defer e.lock.Unlock()
//		e.connections[fd] = conn
//		if len(e.connections)%100 == 0 {
//			log.Printf("total number of connections: %v", len(e.connections))
//		}
//		return nil
//	}
//
//	func (e *epoll) Remove(conn net.Conn) error {
//		fd := socketFD(conn)
//		err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
//		if err != nil {
//			return err
//		}
//		e.lock.Lock()
//		defer e.lock.Unlock()
//		delete(e.connections, fd)
//		if len(e.connections)%100 == 0 {
//			log.Printf("total number of connections: %v", len(e.connections))
//		}
//		return nil
//	}
//
//	func (e *epoll) Wait() ([]net.Conn, error) {
//		events := make([]unix.EpollEvent, 100)
//		n, err := unix.EpollWait(e.fd, events, 100)
//		if err != nil {
//			return nil, err
//		}
//		e.lock.RLock()
//		defer e.lock.RUnlock()
//		var connections []net.Conn
//		for i := 0; i < n; i++ {
//			conn := e.connections[int(events[i].Fd)]
//			connections = append(connections, conn)
//		}
//		return connections, nil
//	}
func socketFD(conn net.Conn) int {
	//tls := reflect.TypeOf(conn.UnderlyingConn()) == reflect.TypeOf(&tls.Conn{})
	// Extract the file descriptor associated with the connection
	//connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	//if tls {
	//	tcpConn = reflect.Indirect(tcpConn.Elem())
	//}
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

// ，这类服务器可能应用于消息推送、IOT、页游等场景，追求的是大量连接，并发量相对不大的场景。还有一类场景是服务器的连接数不多，几十几百，最多几千的TCP连接，比如公司内的服务之间的调用等，？
//
//测试区分两个场景: I/O密集型和计算密集型。I/O密集型的服务比如文件的读取、数据库的访问，远程服务的调用等等，计算密集型的访问比如区块链的挖矿、算法的计算、
//
//类似redis这样的基于内存的数据处理服务等等(当然redis还是memory bound类型的服务)。
//
//我们通过time.Sleep让goroutine休眠来模拟I/O密集型的服务，实际goroutine休眠和真正的I/O密集型的服务还是有区别的，
//
//虽然它们都有一定的耗时，goroutine在等待的过程中会休眠，但是I/O密集型还有大量的I/O访问，比如磁盘、网络等等。出于方便测试的目的，我们还是使用time.Sleep来模拟，主要测试goroutine在休眠一段时间后对性能的影响。
//
//计算密集型的访问我们采用挖矿算法，通过计算hash值，满足一定的挖矿难度让CPU进行大量的计算动作。
//
//测试分别采用并发连接数为 5000、2000、1000、500、200、100，测试对应的吞吐率和延迟。
//
//测试使用多epoller的方式实现的服务器和goroutine-per-connection实现的服务器。因为连接数少，我们可以采用goroutine-per-connection的方式。

// I/O密集型服务是指需要处理大量输入输出(I/O)操作的服务。这些服务通常需要从磁盘、网络或其他存储设备中读取数据，并将数据写入到其他设备中。例如，Web服务器、数据库服务器和邮件服务器都属于I/O密集型服务。
//
//计算密集型服务是指需要进行大量计算操作的服务。这些服务通常需要处理大量的数据，并使用复杂的算法来分析和转换数据。例如，机器学习、图像识别和自然语言处理等服务都属于计算密集型服务。
//
//与I/O密集型服务不同，计算密集型服务的响应时间通常较短，因为它们可以在短时间内完成大量的计算操作。但是，由于它们需要处理大量的数据，因此它们通常需要更高的硬件资源和更复杂的软件架构来实现高效的性能。
