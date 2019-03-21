package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./log/testlog.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("json marshal failed err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a debug logs")
	logs.Info("this is a info logs")
	logs.Warn("this is a waring logs")

}
