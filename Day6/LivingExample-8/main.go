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

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)
	fmt.Printf("sring val:%d\n", val.Elem().Int())
}

func main() {
	var b = 1
	fmt.Println(b)
	testInt(&b)
}
