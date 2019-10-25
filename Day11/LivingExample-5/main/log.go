/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2019-03-12 16:07:14
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-03-12 16:07:14
 */
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
	//return convertLogLevel(appConfig.LogLevel)
	return logs.LevelDebug
}

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appConfig.LogPath
	config["level"] = convertLogLevel

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("init marshal failed,err", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
