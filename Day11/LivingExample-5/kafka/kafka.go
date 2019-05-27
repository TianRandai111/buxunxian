package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type fnwTest struct {
	Name    string
	Message string `json:"message"`
	Type    string `json:"type"`
}

var (
	client sarama.SyncProducer
	err    error
)

func Initkafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//生产者
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka  close,err:", err)
		return
	}

	logs.Debug("init kafka succ")
	return
}

func SendToKafka(data, topic string) (err error) {
	//创建消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	// a := fnwTest{
	// 	Name:    "fnw-dev",
	// 	Message: "this is a good test,hello chaoge!!",
	// 	Type:    "test-test-test",
	// }

	a := data
	// b, err := json.Marshal(a)

	// if err != nil {
	// 	fmt.Println("json.error", err)
	// 	return
	// }

	msg.Value = sarama.StringEncoder(a)
	//发送消息
	_, _, err = client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	//logs.Debug("pid:%v offset:%v\n,", pid, offset, topic)
	return
}
