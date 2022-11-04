package utils

import (
	"encoding/json"
	"github.com/lemuzhi/zinx/ziface"
	"io/ioutil"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数是可以通过zinx.json由用户进行配置
*/

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer //当前zinx全局的Server对象
	Host      string         //当前服务器主机监听的IP
	TcpPort   int            //当前服务器主机监听的端口号
	Name      string         //当前服务器的名称

	/*
		zinx
	*/
	Version          string //当前Zinx的版本号
	MaxConn          int    //当前服务器主机允许的最大链接数
	MaxPackageSize   uint32 //当前Zinx框架数据包的最大值
	WorkerPoolSize   uint32 //当前业务工作Worker池的Goroutine数量
	MaxWorkerTaskLen uint32 //Zinx框架允许用户最多开辟多少个Worker（限定条件）
}

// GlobalObject 定义一个全局对外的GlobalObj
var GlobalObject *GlobalObj

// 提供一个init方法，初始化当前的GlobalObject
func init() {
	//如果配置文件没有加载，使用默认值，用户也可以自己进行相应配置
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.8",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,   //worker工作池的队列个数
		MaxWorkerTaskLen: 1024, //每个worker对应的消息队列的任务数量最大值
	}

	//应该尝试从conf/zinx/json去加载一些用户自定义的参数
	GlobalObject.Reload()
}

// 从zinx.json去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将json文件的数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
