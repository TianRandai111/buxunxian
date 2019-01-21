package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
	"http://www.jd.com",
}

func main() {
	for _, v := range url {
		//这是超时时间
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		//Head head请求
		resp, err := c.Head(v)
		if err != nil {
			fmt.Println("http,head error:", v, err)
			continue
		}
		//Status获取状态码
		fmt.Printf("head succ, status :%v\n", resp.Status)
	}
}
