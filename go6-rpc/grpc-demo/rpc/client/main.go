package main

import (
	"fmt"
	"grpc-demo/rpc/service"
	"net/rpc"
)

// var
// var var1 int = "string"
// var _ Service = &HelloService{}
// int(60)  ---> (int)(<type>)
// 我们声明了一个空指针,强制把这个指针转换成了一个 *HelloService
var _ service.Service = (*HelloServiceClient)(nil)

// NewHelloServiceClient 客户端构造函数
func NewHelloServiceClient(network, address string) (service.Service, error) {
	// 首先是通过rpc.Dial 拨号RPC服务, 建立连接。
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

// Hello 对于RPC客户端来说,我们就需要包装客户端的调用
func (c *HelloServiceClient) Hello(name string, resp *string) error {
	// 然后通过client.Call调用具体的RPC方法
	// 在调用client.Call时:
	// 		第一个参数是用点号链接的RPC服务名字和方法名字，
	// 		第二个参数是 请求参数
	//      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值
	return c.client.Call(service.Name+".Hello", name, resp)
}

func main() {
	// 初始化客户端实例
	client, err := NewHelloServiceClient("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	var resp string
	err = client.Hello("alice", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
