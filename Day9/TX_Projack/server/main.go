package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/TianRandai111/buxunxian/Day9/TX_Projack/common/message"
)

func readPkg(conn net.Conn) (mes message.LoginMes, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")

	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Println("读到的buf=", buf[:4])

	//根据buf[:4]的长度，转换成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容,conn从pkgLen中读取内容到buf中
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) err:", err)
		return
	}

	//把pkgLen反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	//循环读取客户端信息
	for {
		//这里将读取的数据包，直接防撞城一个函数readPkg(),返回Message,Err

	}
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen error=", err)
		return
	}
	//一但监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("连接成功")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//连接成功，启动一个协程和客户端端保持通讯
		go process(conn)
	}
}
