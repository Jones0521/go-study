package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Service interface {
	Hello(name string, resp *string) error
}

// var
var var2 Service = &HelloService{}

type HelloService struct {
}

// Hello 业务场景
// 该函数需要被客户端调用
// name ---->
// <---- hello, name
// 改成符合RPC规范的 函数签名
// 1. 第一个参数 request, interface{}
// 2. 第二次参数是一个响应 reponse，interface{}, 必须是一个指针
// 3. 返回errot
func (h *HelloService) Hello(name string, resp *string) error {
	//return fmt.Sprintf("hello, %s", name)
	*resp = fmt.Sprintf("hello, %s", name)
	return nil
}
func main() {
	// 发送要提供的服务注册给RPC框架
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}
	// 监听socket
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 前面都是tcp的知识, 到这个RPC就接管了
		// 因此 你可以认为 rpc 帮我们封装消息到函数调用的这个逻辑,
		// 提升了工作效率, 逻辑比较简洁，可以看看他代码
		go rpc.ServeConn(conn)
	}

}
