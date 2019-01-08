package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/TianRandai111/buxunxian/Day9/TX_Projack/common/message"
)

var network_address string

func login(userid int, pwd string) (err error) {
	//定义协议
	// fmt.Printf("userid=%d,pwd=%s\n", userid, pwd)
	// return nil
	//1.链接服务器
	network_address = "10.4.83.196:8889"
	conn, err := net.Dial("tcp", network_address)
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userid
	loginMes.UserPwd = userpwd

	//4.将loginmes学历恶化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("data json.Marshal err=", err)
		return
	}
	// 5.data赋值给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" mes json.Marshal err=", err)
		return
	}

	//7. 到这个时候 data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器

	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkglen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write error=", err)
		return
	}
	//fmt.Printf("客户端，发送消息长度成功=%d,内容为=%s\n", len(data), data)
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err", err)
		return
	}

	//这里还需要处理服务器端返回的消息

	return
}
