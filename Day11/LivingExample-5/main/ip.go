package main

import (
	"fmt"
	"net"
)

var (
	localIP []string
)

func InitIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("获取本机IP失败"))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = append(localIP, ipnet.IP.String())
			}
		}
	}
	fmt.Println(localIP)
}
