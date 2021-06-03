package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-study/consul/pb"
	"google.golang.org/grpc"
	"net"
)
type Children struct {

}

func (children *Children)SayHello(ctx context.Context, p *pb.Person)(*pb.Person, error){
	p.Name = "hello " + p.Name
	return p, nil
}



func main() {
	//把grpc服务，注册到consul上
	// 1. 初始化consul配置
	consulConfig := api.DefaultConfig()
	// 2. 创建consul对象
	consulClient,err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.newClient faild", err)
		return
	}
	// 3. 告诉consul即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID : "bj38",
		Tags: []string{"grpc", "consul"},
		Name: "grpc And consul",
		Address: "127.0.0.1",
		Port : 8800,
		Check: &api.AgentServiceCheck{
			CheckID: "consul grpc test",
			TCP : "127.0.0.1:8800",
			Timeout: "5s",
			Interval: "5s",
		},
	}
	//注册服务到consul上
	consulClient.Agent().ServiceRegister(&reg)

	//以下为grpc远程服务调用
	//1.初始化grpc对象
	grpcServer := grpc.NewServer()
	//2.注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	//3.设置监听，设置IP Port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务启动...")
	//4.启动服务
	grpcServer.Serve(listener)

}
