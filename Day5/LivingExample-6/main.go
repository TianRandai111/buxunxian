package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"名字"`
	Age   int    `json:"年龄"`
	Score int    `json:"成绩"`
}

func main() {
	var stu Student = Student{
		Name:  "步荀仙",
		Age:   18,
		Score: 80,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(string(data))
}
