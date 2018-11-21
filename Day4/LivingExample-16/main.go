package main

import (
	"fmt"
)

func testModifyString() {
	s := "你好hello world"
	s1 := []rune(s)
	s1[1] = '0'
	str := string(s1)
	fmt.Println(s1)
	fmt.Println(str)

}

func main() {
	testModifyString()
}
