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

func testStruct() (ret string, err error) {
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
		err = fmt.Errorf("json.marshal failed err%v", err)
		return
	}
	ret = string(data)
	return
}
func test() {
	data, err := testStruct()
	if err != nil {
		fmt.Println("test struct failed,", err)
	}
	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Println("Unmarshal failed,", err)
		return
	}
	fmt.Println(user1)

}

func testMap() (ret string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["姓"] = "步"
	m["名"] = "荀仙"
	m["年龄"] = "100"
	data, err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("json.marshal failed err%v", err)
		return
	}
	ret = string(data)
	return
}

func test2() {
	data, err := testMap()
	if err != nil {
		fmt.Println("test map failed,", err)
		return
	}
	var m map[string]interface{}
	m = make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Unmarshal failed,", err)
		return
	}
	fmt.Println(m)
}
func main() {
	test()
	test2()
}
