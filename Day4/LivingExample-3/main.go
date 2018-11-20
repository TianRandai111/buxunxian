package main

import (
	"errors"
)

func initConfilg() (err error) {
	return errors.New("你出错了")
}

func test() {
	//捕获异常
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// b := 0
	// a := 100 / b
	// fmt.Println(a)
	err := initConfilg()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	for index := 0; index < 5; index++ {
		test()
	}
}
