package main

import (
	"context"
	"fmt"
	"time"

	etcd_client "go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()
	cli.Put(context.Background(), "/logagent/conf/", "123456789")

	for {
		resp := cli.Watch(context.Background(), "/logagent/conf/")
		for wresp := range resp {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}
