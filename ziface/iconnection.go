package ziface

import "net"

// IConnection 定义链接模块的抽象层
type IConnection interface {
	// Start 启动链接 让当前的链接准备开始工作
	Start()

	// Stop 停止链接 结束当前链接的工作
	Stop()

	// GetTCPConnection 获取当前链接的绑定socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前链接模块的链接ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端的TCP状态 IP port
	RemoteAddr() net.Addr

	// SendMsg 发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error

	// SetProperty 设置链接属性
	SetProperty(key string, value interface{})

	// GetProperty 获取链接属性
	GetProperty(key string) (interface{}, error)

	// RemoveProperty 移除链接属性
	RemoveProperty(key string)
}

// HandleFunc 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
