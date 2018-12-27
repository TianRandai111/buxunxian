package main

import "fmt"

type student struct {
	name string
}

func IntChan() {
	var intChan chan int
	intChan = make(chan int, 10)
	intChan <- 10
	a := <-intChan
	fmt.Println("intChan=", a)
}

func MapChan() {
	var mapChan chan map[int]string
	mapChan = make(chan map[int]string, 10)
	m := make(map[int]string)
	v := make(map[int]string)
	m[0] = "001"
	m[1] = "002"
	mapChan <- m
	v = <-mapChan
	fmt.Println("MapChan=", v)
}

func StructChan() {
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)
	stu := student{
		name: "buxunxian",
	}
	stuChan <- &stu
	stu1 := <-stuChan
	fmt.Println("stu01=", stu1)

	var stu2 *student
	stu2, ok := stu1.(*student)
	if !ok {
		fmt.Println("can not error")
	}
	fmt.Println("stu02=", stu2)
}

func main() {
	IntChan()
	MapChan()
	StructChan()
}
