 package main

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

const (
	EtcdKey = "/oldboy/backend/logagent/config/127.0.0.1"
)

type LogConf struct {
	Path string `json:"path"`
	Topic string `json:"topic"`
}

func main(){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()
	var logConfArr = []LogConf{}
	logConfArr = append(
		logConfArr,
		LogConf{
			Path: "/var/log/nginx/access.log",
			Topic: "nginx_log",
		},
	)

	logConfArr = append(
		logConfArr,
		LogConf{
			Path: "/var/log/nginx/error.log",
			Topic: "nginx_log_err",
		},
	)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx,EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
