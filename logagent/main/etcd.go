package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go-study/logagent/tailf"
	"golang.org/x/net/context"
	"strings"
	"time"
)
type EtcdClient struct {
	client *clientv3.Client
	keys []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string)(collectConf []tailf.CollectConf, err error){
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
		fmt.Println(etcdKey)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			continue
		}
		cancel()
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:%v", err)
					continue
				}
				logs.Debug("log config is %v", collectConf)
			}
		}
	}

	initEtcdWatcher()
	return
}

func initEtcdWatcher()  {
	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}

func watchKey(key string){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	for {
		rch := cli.Watch(context.Background(), key)
		var collectConf []tailf.CollectConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v", err)
						getConfSucc = false
						continue
					}
				}
				//fmt.Printf("%s %q: %q\n",  ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				tailf.UpdateConfig(collectConf)
			}
		}
	}
}
