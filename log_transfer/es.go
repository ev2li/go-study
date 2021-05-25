package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	elastic "gopkg.in/olivere/elastic.v2"
)

type LogMessaga struct {
	App string
	Topic string
	Message string
}

var (
	esClient *elastic.Client
)

func initEs(esaddr string)(err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esaddr))
	if err != nil {
		fmt.Println("Connect es err:", err)
		return
	}
	esClient = client
	logs.Debug("Connect es success")

	/*
	fmt.Println("Insert success")*/
	return
}

func sendToEs(topic string, data []byte)(err error){
	msg := &LogMessaga{}
	msg.Topic = topic
	msg.Message = string(data)

	_, err = esClient.Index().
		Index(topic).
		Type(topic).
		//Id("1").
		BodyJson(msg).
		Do()
	if err != nil{
		panic(err)
		return
	}
	return
}