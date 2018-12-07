package main

import (
	"fmt"
)

//接口的定义
type Test interface {
	Print()
	Sleep()
}
type Student struct {
	name  string
	age   int
	score int
}
type Popel struct {
	name string
	age  int
}

//多态
func (p *Student) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age", p.age)
	fmt.Println("score", p.score)
}

func (p *Student) Sleep() {
	fmt.Println("这是student")
}
func (p *Popel) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age", p.age)
}
func (p *Popel) Sleep() {
	fmt.Println("这是Popel")
}
func main() {
	var t Test
	var stu Student = Student{
		name:  "stu1",
		age:   20,
		score: 200,
	}
	t = &stu
	t.Print()
	t.Sleep()

	var p Test
	var stu1 Popel = Popel{
		name: "步荀仙",
		age:  200,
	}
	p = &stu1
	p.Print()
	p.Sleep()
}
