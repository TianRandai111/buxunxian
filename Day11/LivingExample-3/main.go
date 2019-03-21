package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./config")
	if err != nil {
		fmt.Println("new config file falied, err : ", err)
		return
	}
	ip := conf.String("server::listen_ip")

	fmt.Println("IP == ", ip)

	port, err := conf.Int("server::listen_port")
	if err != nil {
		fmt.Println("new config file falied Server::Listen_port, err : ", err)
		return
	}
	fmt.Println("port == ", port)

	logs_level := conf.String("logs::log_level")

	fmt.Println("level == ", logs_level)

	logs_path := conf.String("logs::log_path")
	fmt.Println("path == ", logs_path)

	collect_log_path := conf.String("collect::log_path")
	fmt.Println("collect_log_path=", collect_log_path)

	topic := conf.String("collect::topic")
	fmt.Println("topic=", topic)
}
