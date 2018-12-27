package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func testStruct() {
	user1 := &User{
		UserName: "荀仙",
		NickName: "步",
		Age:      100,
		Birthday: "1500",
		Sex:      "男",
		Email:    "buxunxian@xiandao.com",
		Phone:    "110",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed err%v", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["姓"] = "步"
	m["名"] = "荀仙"
	m["年龄"] = "100"
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("json.marshal failed err%v", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func testSliep() {
	var s []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["姓"] = "步"
	m["名"] = "荀仙"
	m["年龄"] = "100"
	s = append(s, m)

	m = make(map[string]interface{})
	m["姓"] = "天"
	m["名"] = "然呆"
	m["年龄"] = "10"
	s = append(s, m)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json.marshal failed err%v", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSliep()
}
