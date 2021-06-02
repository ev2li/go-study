package main

import (
	"fmt"
	"go-study/micro_services/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	//1.连接grpc服务
	conn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc Dial failed,",err)
	}
	defer conn.Close()

	//2.初始化grpc客户端
	client := pb.NewSayNameClient(conn)
	teacher := &pb.Teacher{
		Name: "zhangsan",
		Age: 30,
	}
	//3.调用远程服务
	t, err:= client.SayHello(context.TODO(),teacher)
	fmt.Println(t, err)
}
