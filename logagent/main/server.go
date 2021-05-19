package main

import (
	"github.com/astaxie/beego/logs"
	"go-study/logagent/kafka"
	"go-study/logagent/tailf"
	"time"
)

func serverRun() (err error) {
	for{
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg)(err error)  {
/*	fmt.Printf("read msg:%s, topic:%s", msg.Msg, msg.Topic)
	fmt.Println()*/
	kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}