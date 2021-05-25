package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"strings"
	"time"
)
type EtcdClient struct {
	client *clientv3.Client
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string)(err error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}
	etcdClient = &EtcdClient{
		client: cli,
	}


	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}
	for _,ip := range localIpArray{
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			continue
		}
		cancel()
		for k, v := range resp.Kvs {
			fmt.Println(k, v)
		}
	}

	return
}
