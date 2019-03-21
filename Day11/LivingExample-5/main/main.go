package main

import (
	"fmt"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/tailf"
	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "../conf/config"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("load conf succ , config : %v appConfig", appConfig)

	err = tailf.InitTail(appConfig.CollectConf, appConfig.Chan_Size)
	if err != nil {
		logs.Error("init tail failed err : %v ", err)
		return
	}

	logs.Debug("initialize succ")
	/* 	go func() {
		var count int
		for {
			count++
			logs.Debug("test fo r logger %d", count)
			time.Sleep(time.Millisecond * 100)
		}
	}() */
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed,err:%v", err)
		return
	}
	logs.Info("Program exited")
}
