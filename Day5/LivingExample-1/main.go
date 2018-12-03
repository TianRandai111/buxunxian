package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var stu Student
	stu.Name = "步荀仙"
	stu.Age = 731
	stu.Score = 99.9

	var stu1 *Student = &Student{
		Age:   20,
		Name:  "天然呆",
		Score: 99.9,
	}

	var stu2 = Student{
		Age:   20,
		Name:  "寻魔",
		Score: 99.9,
	}

	fmt.Println(stu)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Age:%p\n", &stu.Age)
	fmt.Printf("Score:%p\n", &stu.Score)
	fmt.Println(*stu1)
	fmt.Println(stu2)
}
