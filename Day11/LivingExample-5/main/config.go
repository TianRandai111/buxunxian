package main

import (
	"errors"
	"fmt"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/tailf"
	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config
)

type Config struct {
	LogLevel    string
	LogPath     string
	CollectConf []tailf.CollectConf
	Chan_Size   int
	KafkaAddr   string
	Etcdaddr    string
	etcdKey     string
}

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	// fmt.Println(cc.LogPath)
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid conllect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	// fmt.Println("---")
	// fmt.Println(cc.Topic)
	// fmt.Println("---")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid conllect::topic")
		return
	}

	appConfig.CollectConf = append(appConfig.CollectConf, cc)
	return
}

func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed ,err", err)
		return
	}

	appConfig = &Config{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}

	appConfig.Chan_Size, err = conf.Int("collect::chan_size")
	if err == nil {
		appConfig.Chan_Size = 100
	}

	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogLevel = "../logs"
	}

	appConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}
	appConfig.Etcdaddr = conf.String("etcd::addr")
	if len(appConfig.Etcdaddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
		return
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		err = fmt.Errorf("invalid etcd key")
		return
	}

	err = loadCollectConf(conf)

	return
}
