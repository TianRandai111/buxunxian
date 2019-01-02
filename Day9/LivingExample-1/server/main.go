package main

import (
	"fmt"
	"net"
)

//网络socket开发是，net包含有我们需要的所有函数和方法。

func main() {
	fmt.Println("服务器开始监听...")
	List, err := net.Listen("tcp", "0.0.0.0:8888")
	//1.tcp表示使用的网络协议是tcp
	//2.0.0.0.0:8888表示本地监听8888端口

	if err != nil {
		fmt.Println("network listen failed:", err)
		return
	}

	defer List.Close() //延时关闭
	//循环等待客户端连接
	for {
		//等待链接
		fmt.Println("等待客户端连接")
		conn, err := List.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept () suc con=%v\n", conn)
		}
		//启动一个协程，微客户端服务

	}

	fmt.Printf("listen=%v\n", List)
}
