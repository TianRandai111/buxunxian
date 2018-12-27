package main

import (
	"flag"
	"fmt"
)

func main() {
	var confPath string
	var loglevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&loglevel, "d", 0, "please input log level")

	//使之前的flag.xxx生效
	flag.Parse()

	fmt.Println("path:", confPath)
	fmt.Println("log level", loglevel)
}
