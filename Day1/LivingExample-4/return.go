package main

import (
	"fmt"
)

func add(a int, b int) /*定义两个变量来接收值*/ (int, int) /*这里是返回值的类型*/ {
	c := a + b
	d := a * b
	return c, d //这里返回的两个变量的值
}
func main() {
	newc, newd := add(1, 2) //这里将1和2传送到add的函数中，用newc和newd来接受add的返回值
	fmt.Println("newc=", newc, "newc=", newd)
}
