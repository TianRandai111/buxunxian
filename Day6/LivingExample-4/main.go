package main

import "fmt"

type Student struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	//b += 3
	fmt.Println(b)
}

func just(tem ...interface{}) {
	for i, v := range tem {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool, value is %v\n", i, v)
		case int, int64, int32:
			fmt.Printf("%d params is int, value is %v\n", i, v)
		case float32, float64:
			fmt.Printf("%d params is float, value is %v\n", i, v)
		case string:
			fmt.Printf("%d params is string, value is %v\n", i, v)
		case Student:
			fmt.Printf("%d params is Student, value is %v\n", i, v)
		case *Student:
			fmt.Printf("%d params is *Student, value is %v\n", i, v)
		}
	}
}

func main() {
	var b Student
	Test(b)
	a := Student{
		Name: "步荀仙",
		Sex:  "男",
	}
	just(1, true, &a, "This is test", a, 9.9)
}
