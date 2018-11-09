package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Printf("年:%02d\n", time.Now().Year())
	fmt.Printf("月:%02d\n", time.Now().Month())
	fmt.Printf("日:%02d\n", time.Now().Day())

	fmt.Printf("时:%02d\n", time.Now().Hour())
	fmt.Printf("分:%02d\n", time.Now().Minute())
	fmt.Printf("秒:%02d\n", time.Now().Second())
	fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
