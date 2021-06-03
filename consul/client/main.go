package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-study/consul/pb"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	//1.初始化consul配置
	consulConfig := api.DefaultConfig()
	//2.创建consul对象 (可以重新去指定consul属性, IP/Port，也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	//3.服务发现，从consul上获取健康的服务
	services,_, err := consulClient.Health().Service("grpc And consul", "grpc", true, nil)
	if err != nil {
		fmt.Println("consul health service失败!")
		return
	}
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	//=================以下是客户端
	//1.连接服务
	grpcConn,_ := grpc.Dial(addr, grpc.WithInsecure())

	//2.初始化grpc客户端
	grpcClient := pb.NewHelloClient(grpcConn)

	var persoon pb.Person
	persoon.Name = "Andy"
	persoon.Age = 20
	//3.调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &persoon)
	fmt.Println(p, err)
}
