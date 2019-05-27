package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr string
	ESAddr    string
	LogPath   string
	LogLevel  string
	Topic     string
}

var (
	logConfig *LogConfig
)

func initConfig(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed ,err", err)
		return
	}
	logConfig = &LogConfig{}

	logConfig.LogLevel = conf.String("logs::log_level")
	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}

	logConfig.LogPath = conf.String("logs::log_path")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogLevel = "../logs"
	}

	logConfig.KafkaAddr = conf.String("kafka::kafka_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	logConfig.Topic = conf.String("kafka::topic")
	if len(logConfig.Topic) == 0 {
		err = fmt.Errorf("invalid kafka topic")
		return
	}

	logConfig.ESAddr = conf.String("ES::ES_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}
	return
}
