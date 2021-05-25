package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr string
	EsAddr string
	LogPath string
	LogLevel string
	Topic string
}
var (
	logConfig *LogConfig
)

func initConfig(confType, filename string)(err error){
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("new config failed, err:%v", err)
		return
	}
	logConfig = &LogConfig{}
	logConfig.LogLevel = conf.String("logs::log_level")

	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}

	logConfig.LogPath = conf.String("logs::log_path")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "./logs"
	}

	logConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("Invalid kafka addr")
		return
	}

	logConfig.Topic = conf.String("kafka::topic")
	if len(logConfig.Topic) == 0 {
		err = fmt.Errorf("Invalid kafka topic")
		return
	}

	logConfig.EsAddr = conf.String("es::addr")
	if len(logConfig.EsAddr) == 0  {
		err = fmt.Errorf("Invalid es addr")
		return
	}
	return
}