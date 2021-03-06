package main

import (
	"time"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/kafka"
	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/tailf"
	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed , err :%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	logs.Debug("read msg:%s,topic:%s", msg.Msg, msg.Topic)
	//fmt.Printf("read msg:%s,topic:%s", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
