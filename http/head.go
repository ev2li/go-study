package main

import (
	"net/http"
	"net"
	"time"
	"fmt"
)

var url = []string{
		"http://www.baidu.com",
		"http://taobao.com",
}

func main() {
	test301()
}

func test301()  {
	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}
		fmt.Printf("head success, status :%v\n", resp.Status)
	}
}