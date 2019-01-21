package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/TianRandai111/buxunxian/Day9/TX_Projack/common/message"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")

	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pakg header error")
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
func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方
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

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Write error=", err)
		return
	}
	return
}
