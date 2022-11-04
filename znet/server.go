package znet

import (
	"fmt"
	"github.com/lemuzhi/zinx/utils"
	"github.com/lemuzhi/zinx/ziface"
	"net"
)

// Server IServer接口实现，定义一个服务器模块
type Server struct {
	Name      string //服务器名称
	IPVersion string //ip版本
	IP        string //ip地址
	Port      int    //监听端口
	//当前 server的消息管理模块，用来绑定MsgID和对应的处理业务API关系
	MsgHandler ziface.IMsgHandle

	//该server的链接管理器
	ConnMgr ziface.IConnManage
}

func (s *Server) Start() {
	fmt.Printf(`
	欢迎使用Zinx系统!
	Server name : %s
	Server Listenner at IP:%s:%d`, s.Name, s.IP, s.Port)

	go func() {
		//0、开启消息队列及Worker工作池
		s.MsgHandler.StartWorkerPool()

		//1、获取一个TCP的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("解决tcp地址错误:", err)
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

			//设置最大链接个数的判断，如果超过最大链接数量，那么关闭新的链接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//TODO 给客户端响应一个超出最大链接的错误包
				fmt.Println("Too Many Connections MaxConn = ", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//将处理新链接的业务方法和conn进行绑定，得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			//启动当前的链接业务处理
			go dealConn.Start()
		}
	}()
}

// Stop 停止服务器
func (s *Server) Stop() {
	//将一些服务器的资源，状态或者一些已经开辟的链接信息，进行停止或者回收
	fmt.Println("[STOP] Zinx server name ", s.Name)
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态，避免Serve结束，导致Start提前结束
	select {}
}

// AddRouter 路由功能： 给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouer(msgID, router)
	fmt.Println("Add Router Success!!!")
}

func (s *Server) GetConnMgr() ziface.IConnManage {
	return s.ConnMgr
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}
