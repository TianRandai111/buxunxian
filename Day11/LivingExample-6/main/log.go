package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return convertLogLevel(Server.Logs_Conf.Log_Level)
}

func Init_Logs() {
	config := make(map[string]interface{})
	config["filename"] = Server.Logs_Conf.Log_Path
	config["level"] = Server.Logs_Conf.Log_Level

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("init marshal failed :", err)
		logs.Debug("init marshal failed :", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
