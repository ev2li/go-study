package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go-study/logagent/kafka"
	"go-study/logagent/tailf"
)

func main() {
	filename := "../../conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err : %v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err : %v\n", err)
		panic("load logger failed")
		return
	}
	logs.Debug("initialize success")
	logs.Debug("load conf success, config:%v", appConfig)

	err = tailf.InitTail(appConfig.CollectConf, appConfig.ChanSize)
	if err != nil {
		logs.Error("init tail faild, err : %v", err)
		return
	}

	err = kafka.InitKafka(appConfig.KafkaAddr)
	if err != nil {
		logs.Error("init kafka faild, err : %v", err)
		return
	}

	logs.Debug("initialize all success")
/*	go func() {
		var count int
		for {
			count++
			logs.Debug("test for logger %d", count)
			time.Sleep(time.Second)
		}
	}()*/
	err = serverRun()
	if err != nil {
		logs.Error("serverRun faild, err : %v", err)
		return
	}
	logs.Info("program exited")
}