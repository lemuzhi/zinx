package ziface

// 定义一个服务器模块的抽象层接口
type IServer interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//运行服务
	Serve()

	//路由功能： 给当前的服务注册一个路由方法，供客户端的链接处理使用
	AddRouter(msgID uint32, router IRouter)
}
