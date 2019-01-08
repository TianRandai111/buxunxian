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
	_, err = conn.Do("HMSet", "user01", "name", "buxunxian", "age", 18)
	if err != nil {
		fmt.Println("数据插入失败：", err)
		return
	}

	fmt.Println("数据插入成功")

	//strings转换多个值
	data, err := redis.Strings(conn.Do("HMGet", "user01", "name", "age"))
	if err != nil {
		fmt.Println("get获取数据失败", err)
		return
	}

	for _, i := range data {
		fmt.Println(i)
	}

	//返回的data是interface{}
	//因为name对应的值是string,因此我们需要转换

}
