package znet

import (
	"fmt"
	"github.com/lemuzhi/zinx/ziface"
	"net"
)

//IServer接口实现，定义一个服务器模块
type Server struct {
	Name string //服务器名称
	IPVersion string //ip版本
	IP string //ip地址
	Port int //监听端口
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
		fmt.Println("句柄=",addr)
		//2、监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("监听", addr, "错误：", err)
			return
		}
		fmt.Println("Zinx启动成功，", s.Name)
		//3、阻塞等待客户端链接，处理客户端链接业务（读写）
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("套接字失败：", err)
				continue
			}

			//已经与客户端建立链接，做一些业务，做一个最基本的最大512字节长度的回写业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("接收数据出错：", err)
						continue
					}
					//回写功能
					_, err = conn.Write(buf[:cnt])
					if err != nil {
						fmt.Println("回写数据失败：", err)
						continue
					}
				}
			}()
		}
	}()
}


func (s *Server) Stop()  {

}


func (s *Server) Serve()  {
	fmt.Println("开始运行")
	s.Stop()
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