package main

import (
	"fmt"
	"net/rpc"
)

type MyClient struct {
	c *rpc.Client
}

func (this *MyClient) HelloWorld(a string, b *string) error {
	//参数1，参照上面的Interface,RegisterName而来，a:传入参数，b:传出参数
	return  this.c.Call("hello.HelloWorld", a, b)
}

func InitClient(addr string) MyClient {
	conn, _ := rpc.Dial("tcp", addr)
	return MyClient{
		c : conn,
	}
}

func main() {
	//1.用rpc去转接服务器
	//conn, err := rpc.Dial("tcp",  "127.0.0.1:8800")
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	myClient := InitClient("127.0.0.1:8800")
	var resp string
	err := myClient.HelloWorld("杜甫", &resp)
	if err != nil {
		fmt.Println("HelloWorld error", err)
		return
	}

	fmt.Println(resp, err)
}
