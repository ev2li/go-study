package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(kafkaAddr string)(err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		logs.Error("Init kafka producer failed, err:", err)
		return
	}
	logs.Debug("Init kafka success")
	return
}

func SendToKafka(data, topic string)(err error){
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed err:%v data:%v topic:%v ", err, data, topic)
		return
	}
	logs.Debug("send success pid:%v, offset:%v, topic: %v\n", pid, offset, topic)
	return
}