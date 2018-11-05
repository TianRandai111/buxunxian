package main

import (
	"fmt"
)

//字符串反转
func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

//数组
func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
		//append关键字是用来做什么的
	}
	return string(result)
}

func main() {
	reverse := reverse("hello world")
	fmt.Println(reverse)
	reverse = reverse1("hello world")
	fmt.Println(reverse)
}
