package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Set(name string, age int, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func TsetStuct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect strcut")
		return
	}
	//结构体的数量
	num := val.NumField()
	fmt.Println(num)
	//结构体方法的数量
	numm := val.NumMethod()
	fmt.Println(numm)
}

func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}
	TsetStuct(a)
	fmt.Println(a)
}
