package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "/var/log/nginx/access.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location: &tail.SeekInfo{Offset: 0,Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
