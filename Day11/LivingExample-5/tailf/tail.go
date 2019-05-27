package tailf

import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/hpcloud/tail"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail     *tail.Tail
	conf     CollectConf
	status   int
	exitChan chan int
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObjMgr struct {
	tailobjs []*TailObj
	msgChan  chan *TextMsg
	locl     sync.Mutex
}

var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func UPdateConfig(confs []CollectConf) (err error) {
	tailObjMgr.locl.Lock()
	defer tailObjMgr.locl.Lock()
	for _, oneConf := range confs {
		var isRunning = false
		for _, obj := range tailObjMgr.tailobjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
			if isRunning {
				continue
			}
			createNewTask(oneConf)
		}
	}

	var TailObjs []*TailObj
	for _, obj := range tailObjMgr.tailobjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}
		if obj.status == StatusDelete {
			obj.exitChan <- 1
		}
		TailObjs = append(TailObjs, obj)
	}

	tailObjMgr.tailobjs = TailObjs
	return
}

func createNewTask(conf CollectConf) {
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}

	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//location: &tail.SeekInfo{Offset: 0,Whence:2},
		MustExist: false,
		Poll:      true,
	})
	if errTail != nil {
		logs.Error("collect filename[%s] failed,err :%v", conf.LogPath, errTail)
		return
	}
	obj.tail = tails
	tailObjMgr.tailobjs = append(tailObjMgr.tailobjs, obj)

	go readFromTail(obj)
}

func InitTail(conf []CollectConf, chanSize int) (err error) {
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}
	if len(conf) == 0 {
		//err = fmt.Errorf("invalid config for log collect, conf:%v,", conf)
		logs.Error("invalid config for log collect, conf:%v,", conf)
		return
	}

	fmt.Println(chanSize)
	for _, v := range conf {
		createNewTask(v)
		/* 		obj := &TailObj{
		   			conf: v,
		   		}

		   		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
		   			ReOpen: true,
		   			Follow: true,
		   			//location: &tail.SeekInfo{Offset: 0,Whence:2},
		   			MustExist: false,
		   			Poll:      true,
		   		})
		   		if errTail != nil {
		   			logs.Error("invalid config for log collect, conf:%v,err:%v,", conf, errTail)
		   			//err = errTail
		   			return
		   		}
		   		obj.tail = tails
		   		tailObjMgr.tailobjs = append(tailObjMgr.tailobjs, obj)

		   		go readFromTail(obj) */
	}
	return
}

func readFromTail(tailObj *TailObj) {
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg:   line.Text,
				Topic: tailObj.conf.Topic,
			}
			tailObjMgr.msgChan <- textMsg
		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited,conf:%v", tailObj.conf)
			return
		}
	}
}
