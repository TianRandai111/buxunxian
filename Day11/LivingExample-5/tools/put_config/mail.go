package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/tailf"

	etcd_client "go.etcd.io/etcd/clientv3"
)

const (
	EtcdKey = "/logagent/config/10.4.82.60"
)

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"Topic"`
}

func SetLogConfToEtcd() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"10.39.6.202:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	var logCanfArr []tailf.CollectConf
	logCanfArr = append(
		logCanfArr,
		tailf.CollectConf{
			LogPath: "/tmp/nginx/access.log",
			//LogPath: "/var/log/nginx/access.log",
			Topic: "nginx_log",
		},
	)

	logCanfArr = append(
		logCanfArr,
		tailf.CollectConf{
			LogPath: "/tmp/nginx/error.log",
			//LogPath: "/var/log/nginx/error.log",
			Topic: "nginx_log_err",
		},
	)

	data, err := json.Marshal(logCanfArr)
	if err != nil {
		fmt.Println("json falied,", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put falied ,err ", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("put falied ,err ", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
