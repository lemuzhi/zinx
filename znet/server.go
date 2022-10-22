package znet

import (
	"errors"
	"fmt"
	"github.com/lemuzhi/zinx/ziface"
	"net"
)

// Server IServer接口实现，定义一个服务器模块
type Server struct {
	Name string //服务器名称
	IPVersion string //ip版本
	IP string //ip地址
	Port int //监听端口
}

// CallBackToClient 定义当前客户端链接所绑定的HandleAPI（这个hanlde目前写死的，以后优化应该由用户自定义handle方法）
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显的业务
	fmt.Println("Conn Handle CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("To client callback write data error: ", err)
		return errors.New("CallBackToClient error")
	}

	return nil
}

func (s *Server) Start()  {
	fmt.Printf("Server Listenner at IP:%s:%d\n", s.IP, s.Port)

	go func() {
		//1、获取一个TCP的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("解决tcp地址错误:", err)
			return
		}
		//fmt.Println("句柄=",addr)
		//2、监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("监听", addr, "错误：", err)
			return
		}
		fmt.Println("Zinx启动成功，", s.Name)

		var cid uint32
		cid = 0

		//3、阻塞等待客户端链接，处理客户端链接业务（读写）
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("套接字失败：", err)
				continue
			}

			//将处理新链接的业务方法和conn进行绑定，得到我们的链接模块
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			//启动当前的链接业务处理
			go dealConn.Start()
		}
	}()
}


func (s *Server) Stop()  {

}


func (s *Server) Serve()  {
	fmt.Println("启动服务器")
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态，避免Serve结束，导致Start提前结束
	select {

	}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s
}