package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"go-study/logagent/tailf"
)

var (
	appConfig *Config
)
type Config struct {
	LogLevel string
	LogPath string
	ChanSize int
	KafkaAddr string
	CollectConf []tailf.CollectConf

	etcdAddr string
	etcdKey string
}

func loadCollectConf(conf config.Configer)(err error){
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("Invalid collect::log_path")
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("Invalid collect::topic")
	}

	appConfig.CollectConf = append(appConfig.CollectConf, cc)
	return
}

func loadConf(confType, filename string)(err error){
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("new config failed, err:%v", err)
		return
	}
	appConfig = &Config{}
	appConfig.LogLevel = conf.String("logs::log_level")

	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}

	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs"
	}

	appConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("Invalid kafka addr")
		return
	}
	appConfig.ChanSize , err = conf.Int("logs::chan_size")
	if err != nil {
		appConfig.ChanSize = 100
	}

	appConfig.etcdAddr = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0  {
		err = fmt.Errorf("Invalid etcd addr")
		return
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0  {
		err = fmt.Errorf("Invalid etcd configKey")
		return
	}
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	return
}
