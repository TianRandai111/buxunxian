package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		/* Dial: func(redis.Conn, error) {
			return redis.Dial("tcp", "10.39.6.202:6379")
		}, */
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "10.39.6.202:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "name", "步荀仙")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

	conn2 := pool.Get()
	defer conn2.Close()
	_, err = conn2.Do("Set", "name2", "天然呆")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r2, err := redis.String(conn.Do("Get", "name2"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r2=", r2)
}
