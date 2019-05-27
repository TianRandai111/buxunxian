package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	wg sync.WaitGroup
)

func run() (err error) {

	partitionList, err := kafkaClient.client.Partitions(kafkaClient.topic)
	if err != nil {
		fmt.Println("Failed to get the list of partitions", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, errRet := kafkaClient.client.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			err = errRet
			logs.Error("Failed to start kafkaClient.client for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			wg.Add(1)
			for msg := range pc.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				err = sendtToES(kafkaClient.topic, msg.Value)
				if err != nil {
					logs.Warn("send to es failed ,err:%v", err)
				}
			}
			wg.Done()
		}(pc)
	}
	time.Sleep(time.Hour)
	wg.Wait()
	return
}
