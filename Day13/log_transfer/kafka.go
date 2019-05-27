package main

import (
	"strings"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)



var (
	kafkaClient *KafkaClient
)

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
}

func initKafka(kafka, topic string) (err error) {

	kafkaClient = &KafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(kafka, ","), nil)
	if err != nil {
		logs.Error("Failed to get the list of Consumer:", err)
		return
	}
	kafkaClient.client = consumer
	kafkaClient.addr = kafka
	kafkaClient.topic = topic
	return
}
