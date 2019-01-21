package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/TianRandai111/buxunxian/Day9/TX_Projack/common/message"
)

func readPkg(conn net.Conn) (mes message.LoginMes, err error) {

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

func process(conn net.Conn) {
	defer conn.Close()

	//循环读取客户端信息
	for {
		//这里将读取的数据包，直接防撞城一个函数readPkg(),返回Message,Err
		mes, err := readPkg(conn)
		if err != nil {
			fmt.Println("readPkg(conn) err=", err)
		}

		fmt.Println("mes=", mes)

	}
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

//编写一个函数处理登陆 serverProcessMesLogin
func serverProcessMesLogin(conn net.Conn, mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Marshal([]byte(mes.Data), &loginMes) err", err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginMesType

	//在声明一个LoginResMes
	var loginResMes message.LoginRseMes
	//如果用户的ID =100 ，密码登陆=123456 是合法的

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
		loginResMes.Error = "用户登陆陈宫"
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册在使用"
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) failed:", err)
		return
	}
	// 将data复制给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) failed:", err)
		return
	}

	//6. 发送data,我们将其封装到writaPkg函数
	err = writePkg(conn, data)
	return
}

//编写一个ServerProcessMae函数
//功能：根据客户端发送消息的种类的不同，巨鼎调用那个函数来处理

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginRseMesType:
		//处理登陆
		err = serverProcessMesLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册
	default:
		return
	}
	return
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
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出", err)
				return
			} else {
				fmt.Println("listen.Accept err=", err)
				return
			}
			//连接成功，启动一个协程和客户端端保持通讯
			go process(conn)

		}
	}
}
