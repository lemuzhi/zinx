package znet

//IServer接口实现，定义一个服务器模块
type Server struct {
	Name string //服务器名称
	IPVersion string //ip版本
	IP string //ip地址
	Port int //监听端口
}

func (s *Server) Start()  {
	
}


func (s *Server) Stop()  {

}


func (s *Server) Serve()  {

}

func NewServer(name string) *Server {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s
}