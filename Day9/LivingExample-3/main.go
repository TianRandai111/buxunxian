package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "10.39.6.202:6379")
	if err != nil {
		fmt.Println("连接Redis失败", err)
		return
	}
	defer conn.Close()

	//set值
	fmt.Println("连接Redis成功:", conn)
	_, err = conn.Do("HSet", "user01", "name", "buxunxian")
	if err != nil {
		fmt.Println("数据插入失败：", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("数据插入失败：", err)
		return
	}
	fmt.Println("数据插入成功")

	data, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("get获取数据失败", err)
		return
	}

	data1, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("get获取数据失败", err)
		return
	}
	//返回的data是interface{}
	//因为name对应的值是string,因此我们需要转换

	fmt.Println("data1=", data)
	fmt.Println("data1=", data1)

}
