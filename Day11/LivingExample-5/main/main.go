package main

import (
	"fmt"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/kafka"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/tailf"
	"github.com/astaxie/beego/logs"
)

func main() {
	InitIp()

	filename := "../conf/config"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}
	logs.Debug("load conf succ , config : %v appConfig", appConfig)

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	collectConf, err := InitEtcd(appConfig.Etcdaddr, appConfig.etcdKey)
	if err != nil {
		logs.Error("init etcd failed err : %v ", err)
		return
	}

	logs.Debug("initialize etcd succ")

	err = tailf.InitTail(collectConf, appConfig.Chan_Size)
	if err != nil {
		logs.Error("init tail failed err : %v ", err)
		return
	}

	err = kafka.Initkafka(appConfig.KafkaAddr)
	if err != nil {
		logs.Error("init kafka failed err : %v ", err)
		return
	}
	logs.Debug("initialize kafka succ")

	/* 	go func() {
		var count int
		for {
			count++
			logs.Debug("test fo r logger %d", count)
			time.Sleep(time.Millisecond * 100)
		}
	}() */
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed,err:%v", err)
		return
	}
	logs.Info("Program exited")
	logs.Debug("initialize succ")
}
