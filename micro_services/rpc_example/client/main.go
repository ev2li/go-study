package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	//1.用rpc去转接服务器
	//conn, err := rpc.Dial("tcp",  "127.0.0.1:8800")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return

	}
	defer conn.Close()
	//2.调用远程函数
	var reply string
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}

	fmt.Println(reply)
}
