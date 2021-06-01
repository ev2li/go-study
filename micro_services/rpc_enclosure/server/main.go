package main

import (
	"fmt"
	"net"
	"net/rpc"
)
//创建接口，在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

//调用方法时，需要传参，参数应该是实现了HelloWorld方法对象
func RegisterService(i MyInterface) {
	err := rpc.RegisterName("hello", i)
	if err != nil {
		fmt.Println("注册rpc服务失败!", err)
		return
	}
}


type World struct {

}

func (this *World) HelloWorld(name string,resp *string) error {
	*resp = name + " 你好!"
	return nil
}

func main(){
	//1.注册rpc服务，绑定对象方法
     RegisterService(new(World))
	//2.设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Net listen err:", err)
		return
	}

	defer listener.Close()
	fmt.Println("开始监听...")
	//3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()
	//4.绑定服务
	rpc.ServeConn(conn)
}

