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

func insertHead(p **Student) {
	for i := 9; i >= 0; i-- {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.Next = *p
		*p = &stu
	}
	//inse
}

func delNode(p *Student) {
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6" {
			prev.Next = p.Next
			break
		}
		prev = p
		p = p.Next
	}
}

func addNode(p, new *Student) {
	for p != nil {
		if p.Name == "stu7" {
			new.Next = p.Next
			p.Next = new
			break
		}
		p = p.Next
	}
}

func main() {
	var head *Student = &Student{}
	head.Name = "步荀仙"
	head.Age = 99
	head.Score = 100

	insertHead(&head)
	trans(head)
	fmt.Printf("--------------------------------------------------------\n")
	delNode(head)
	trans(head)
	fmt.Printf("--------------------------------------------------------\n")

	var new *Student = new(Student)
	new.Name = "树常青"
	new.Age = 19
	new.Score = 100

	addNode(head, new)
	trans(head)
	fmt.Printf("--------------------------------------------------------\n")
}
