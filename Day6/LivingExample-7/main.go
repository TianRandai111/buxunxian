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

func test(b interface{}) {
	//类型
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v,%T\n", stu, stu)
	}

}

func testInt(b interface{}) {
	fmt.Printf("get value b %d\nget Type b %T\n", b, b)
	val := reflect.ValueOf(b)
	c := val.Int()
	fmt.Printf("get value interface() %d\nget Type interface() %T\n", c, c)
}

func main() {
	var a Student = Student{
		Name:  "步荀仙",
		Age:   18,
		Score: 92,
	}
	test(a)
	testInt(1234)
}
