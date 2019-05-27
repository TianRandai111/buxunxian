package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	err := initConfig("ini", "./conf/config")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(logConfig)

	err = initLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		panic(err)
		return
	}

	logs.Debug("This is test logs")

	err = initKafka(logConfig.KafkaAddr, logConfig.Topic)
	if err != nil {
		logs.Error("init kafka failed,%v", err)
		return
	}
	logs.Debug("This is test kafka")

	err = initES(logConfig.ESAddr)
	if err != nil {
		logs.Error("init es failed,%v", err)
		return
	}
	logs.Debug("This is test es")

	err = run()
	if err != nil {
		logs.Error("init es failed,%v", err)
		return
	}

	logs.Warn("waring,log_transfer is exited")
}
