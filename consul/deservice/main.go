package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-study/consul/pb"
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
	// 3. 注销服务
	consulClient.Agent().ServiceDeregister("bj38")

}
