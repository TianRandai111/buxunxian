package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "10.4.83.10:8888")
	if err != nil {
		fmt.Println("Dial连接失败:", err)
		return
	}
	fmt.Println("Dial 连接成功，conn=", conn)
	//客户端发送数据
	reader := bufio.NewReader(os.Stdin)

	//从终端读取一行用户输入，并准备发送给服务器
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readerstring failed err=:", err)
	}
	value, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("net Write failed err=:", err)
	}
	fmt.Printf("客户单发送了%d字节的数据并退出了", value)
}
