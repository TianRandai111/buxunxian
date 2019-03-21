package main

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/Shopify/sarama"
)

type fnwTest struct {
	Name    string
	Testid  int
	Message string `json:"message"`
	Type    string `json:"type"`
}

//消息写入kafka
func main() {
	//初始化配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//生产者
	client, err := sarama.NewSyncProducer([]string{"10.39.6.202:9092"}, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}

	defer client.Close()
	var a fnwTest
	n := 1
	for {

		//创建消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		a = fnwTest{
			Name:    "fnw-dev",
			Testid:  n,
			Message: "this is a good test,hello chaoge!!",
			Type:    "test-test-test",
		}

		b, err := json.Marshal(a)

		if err != nil {
			fmt.Println("json.error", err)
			return
		}

		msg.Value = sarama.StringEncoder(b)
		//发送消息
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n,", pid, offset)
		time.Sleep(3 * time.Second)
		n++

	}

}
