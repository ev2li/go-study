package main

import (
	"context"
	"fmt"
	"go-study/micro_services/grpc/pb"
	"google.golang.org/grpc"
	"net"
)
// Children 定义类
type Children struct {

}

//SayHello 按接口去绑定类方法
func (child *Children) SayHello(ctx context.Context, teacher *pb.Teacher)(*pb.Teacher, error) {
	teacher.Name += " is sleep!"
	return teacher, nil
}

func main() {
	//1.初始化一个grpc对象
	grpcServer := grpc.NewServer()
	//2.注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))

	//3.设置监听，指定IP,Port
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	//4.启动服务
	grpcServer.Serve(listener)
}
