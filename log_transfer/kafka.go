package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
	"sync"
)

type KafkaClient struct {
	client sarama.Consumer
	addr string
	topic string
	wg sync.WaitGroup
}

var (
	kafkaClient *KafkaClient
)

func initKafka(kafkaAddr, topic string)(err error) {
	kafkaClient = &KafkaClient{}
	consumer, err := sarama.NewConsumer(strings.Split(kafkaAddr, ","), nil)

	if err != nil {
		logs.Error( "Failed to start consumer:%s", err)
		return
	}
	kafkaClient.client = consumer
	kafkaClient.addr = kafkaAddr
	kafkaClient.topic = topic
	return
}
