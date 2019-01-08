package main

import (
	"fmt"
	"net" //网络socket开发是，net包含有我们需要的所有函数和方法。
)

func prossice(conn net.Conn) {
	//这里循环接受客户端发送到的数据
	defer conn.Close() //关闭链接

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等客户端通过Conn发送信息
		//2.如果客户没有wrtie[发送]，那么协程就阻塞在这里
		fmt.Println("服务器在等待客户端" + conn.RemoteAddr().String() + "的输入")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端已退出")
			return
		}
		fmt.Println("客户端传来的多少个字节：", n)

		//3.显示客户端发送的内容到服务器的终端
		fmt.Printf(string(buf[:n]))
	}
}

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
			fmt.Printf("Accept () suc con=%v,客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		//启动一个协程，微客户端服务
		go prossice(conn)
	}

	fmt.Printf("listen=%v\n", List)
}
