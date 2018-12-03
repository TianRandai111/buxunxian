package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	Next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

func main() {
	var head Student
	head.Name = "步荀仙"
	head.Age = 99
	head.Score = 100

	var stu1 Student
	stu1.Name = "天然呆"
	stu1.Age = 27
	stu1.Score = 99

	head.Next = &stu1

	var stu2 Student
	stu2.Name = "寻魔"
	stu2.Age = 20
	stu2.Score = 97

	stu1.Next = &stu2
	trans(&head)

}
