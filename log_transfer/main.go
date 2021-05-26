package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "../../conf/logtransfer.conf"
	err := initConfig("ini", filename)
	if err != nil {
		panic(err)
		return
	}
	err = initLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		panic(err)
		return
	}

	err = initKafka(logConfig.KafkaAddr, logConfig.Topic)
	if err != nil {
		logs.Error("initialize kafka failed, err:%v", err)
		return
	}

	err = initEs(logConfig.EsAddr)
	if err != nil {
		logs.Error("initialize es failed, err:%v", err)
		return
	}

	err = run()
	if err != nil {
		logs.Error("run failed, err:%v", err)
		return
	}

	logs.Warn("waring, log transfer is exited")
}