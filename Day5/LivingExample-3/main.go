package main

import (
	"fmt"
	"math/rand"
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

func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.Next = stu
		tail = stu
	}
}

func main() {
	var head Student
	head.Name = "步荀仙"
	head.Age = 99
	head.Score = 100
	insertTail(&head)
	trans(&head)
}
