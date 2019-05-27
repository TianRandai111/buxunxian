package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/TianRandai111/buxunxian/Day11/LivingExample-5/tailf"
	"go.etcd.io/etcd/mvcc/mvccpb"

	"github.com/astaxie/beego/logs"

	"go.etcd.io/etcd/clientv3"
)

type Etcd_Client struct {
	client *clientv3.Client
	keys   []string
}

var etcdClient *Etcd_Client

var (
	Etcd_Addr string
)

func InitEtcd(addr, key string) (collectConf []tailf.CollectConf, err error) {
	Etcd_Addr = addr
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{Etcd_Addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed,err", err)
		return
	}
	fmt.Println("connect succ")

	etcdClient = &Etcd_Client{
		client: cli,
	}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	//var collectConf []tailf.CollectConf
	for _, ip := range localIP {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed,err:%v", err)
			continue
		}
		cancel()

		logs.Debug("resp from etcd:%v", resp.Kvs)

		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				fmt.Printf("etcd :%v\n", resp)
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed ,err:%v", err)
					continue
				}
			}
			logs.Debug("test logs config is %v", collectConf)
			fmt.Printf("test logs config is %v", collectConf)
		}
	}
	initEtcdWatcher()
	return
}
func initEtcdWatcher() {
	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}

func watchKey(key string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{Etcd_Addr},
		DialTimeout: 5 * time.Second,
	})

	fmt.Printf("%s\n", key)
	if err != nil {
		logs.Error("connect etcd failed,err", err)
		return
	}
	for {
		rch := cli.Watch(context.Background(), key)
		var collectConf []tailf.CollectConf
		var getConfSucc = true
		for wersp := range rch {
			for _, ev := range wersp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config delect'", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key [%s] , unmarshal [%s],err %v", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				tailf.UPdateConfig(collectConf)
			}
		}

	}

}
