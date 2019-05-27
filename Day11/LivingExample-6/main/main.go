package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "../conf/config"
	err := Init_Conf("ini", filename)
	if err != nil {
		fmt.Printf("Config file error :%v\n", err)
	}

	/*  err := Init_Logs()
	if err != nil {
		fmt.Printf("Logs Config error :%v\n", err)
	} */
	for {
		logs.Debug("test fo r logger")
		time.Sleep(time.Millisecond * 100)
	}
	/* 	go func() {
		var count int
		for {
			count++
			logs.Debug("test fo r logger %d", count)
			time.Sleep(time.Millisecond * 100)
		}
	}() */
}
