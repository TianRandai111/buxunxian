package main

import (
	"fmt"
	"os"
)

//定义两个全局变量，一个表示用户ID，一个表示用户密码
var (
	userid  int
	userpwd string
)

func main() {
	var (
		key  int
		loop = true
	)
	for loop {
		//菜单显示
		fmt.Println("----------------------欢迎登陆多人聊天系统----------------------")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("请选择(1-3):")
		fmt.Println("--------------------------------------------------------------")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆中...")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	//根据用户的输入显示新的提示信息
	if key == 1 {
		fmt.Println("请输入用户的ID：")
		fmt.Scanf("%d\n", &userid)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userpwd)
		//登陆函数写在login.go中
		err := login(userid, userpwd)
		if err != nil {
			fmt.Println("登陆失败")
		} else {
			fmt.Println("登陆成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
}
