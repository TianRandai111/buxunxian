package main

import (
	"fmt"
)

func main() {
	var str1 = "hello"
	str2 := "world"
	str3 := str1 + " " + str2
	n := len(str3)
	//字符串的拼装
	fmt.Println(str3)
	fmt.Printf("%s %s\n", str1, str2)
	fmt.Printf("len(str3)=%d\n", n)

	//字符串的截取[切片]
	substr := str3[0:5]
	fmt.Println(substr)

	substr = str3[6:]
	fmt.Println(substr)

}
