package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/config"
)

type LogAgent_Config struct {
	Server_Conf  LogAgent_Server
	Logs_Conf    LogAgent_Logs
	Collect_Conf LogAgent_Collect
	KafKa_Conf   LogAgent_Kafka
}

type LogAgent_Server struct {
	Listen_Ip   string
	Listen_Port int
}

type LogAgent_Logs struct {
	Log_Level string
	Log_Path  string
}

type LogAgent_Collect struct {
	Agent_Path string
	Topic_Name string
	Chan_Size  int
}

type LogAgent_Kafka struct {
	KafKa_Addr string
}

var Server = &LogAgent_Config{}

func Init_Conf(File_Type, File_Name string) (err error) {
	conf, err := config.NewConfig(File_Type, File_Name)
	if err != nil {
		fmt.Println("Load config error：", err)
		logs.Debug("Load config error：", err)
		return
	}

	/* 	var Server LogAgent_Config */

	Server.Server_Conf.Listen_Ip, Server.Server_Conf.Listen_Port, err = load_Server_Conf(conf)
	if err != nil {
		fmt.Println("load Server config error:", err)
		logs.Debug("load Server config error:", err)
		return
	}

	Server.Logs_Conf.Log_Level, Server.Logs_Conf.Log_Path, err = load_Logs_Conf(conf)
	if err != nil {
		fmt.Println("load Logs config error:", err)
		logs.Debug("load Logs config error:", err)
		return
	}

	Server.Collect_Conf.Agent_Path, Server.Collect_Conf.Topic_Name, Server.Collect_Conf.Chan_Size, err = load_Collect_Conf(conf)
	if err != nil {
		fmt.Println("load Collect config error:", err)
		logs.Debug("load Collect config error:", err)
		return
	}

	Server.KafKa_Conf.KafKa_Addr, err = load_Kafka_Conf(conf)
	if err != nil {
		fmt.Println("load Kafka config error:", err)
		logs.Debug("load Kafka config error:", err)
		return
	}
	fmt.Println(Server)
	return

}

func load_Server_Conf(conf config.Configer) (ser_ip string, ser_port int, err error) {
	ser_ip = conf.String("server::listen_ip")
	if len(ser_ip) == 0 {
		fmt.Println("load Server IP is spaces")
		logs.Debug("load Server IP is spaces")
		return
	}

	ser_port, err = conf.Int("server::listen_port")
	if err != nil {
		fmt.Println("load Server Port error", err)
		logs.Debug("load Server Port error", err)
		return
	}
	return
}

func load_Logs_Conf(conf config.Configer) (log_level, log_path string, err error) {
	log_level = conf.String("logs::log_level")
	if len(log_level) == 0 {
		fmt.Println("loading log_level failed")
		logs.Debug("loading log_level failed")
		return
	}

	log_path = conf.String("logs::log_path")
	if len(log_level) == 0 {
		fmt.Println("loading log_path failed")
		logs.Debug("loading log_path failed")
		return
	}
	return
}

func load_Collect_Conf(conf config.Configer) (cc_path, cc_topic string, cc_cz int, err error) {
	cc_path = conf.String("collect::agent_path")
	if len(cc_path) == 0 {
		fmt.Println("loading collect agent path failed")
		logs.Debug("loading collect agent path failed")
		return
	}

	cc_topic = conf.String("collect::topic")
	if len(cc_topic) == 0 {
		fmt.Println("loading collect topic failed")
		logs.Debug("loading collect topic failed")
		return
	}

	cc_cz, err = conf.Int("collect::chan_size")
	if err != nil {
		fmt.Println("load chan size error", err)
		logs.Debug("load chan size error", err)
		return
	}

	return
}

func load_Kafka_Conf(conf config.Configer) (kafka string, err error) {
	kafka = conf.String("kafka::server_addr")
	if len(kafka) == 0 {
		fmt.Println("loading kafka conf failed")
		logs.Debug("loading kafka conf failed")
		return
	}
	return
}
