package main

import (
	"fmt"
)

func testMap() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["abc"] = "efg"
	fmt.Println(a)
}

func testMapTwo() {
	b := make(map[string]map[string]string, 10)
	b["abc1"] = make(map[string]string)
	b["abc1"]["key1"] = "value1"
	b["abc1"]["key2"] = "value2"
	b["abc1"]["key3"] = "value3"
	b["abc1"]["key4"] = "value4"
	b["abc1"]["key5"] = "value5"
	b["abc1"]["key6"] = "value6"
	b["abc2"] = make(map[string]string)
	b["abc2"]["key1"] = "value1"
	b["abc2"]["key2"] = "value2"
	b["abc2"]["key3"] = "value3"
	b["abc2"]["key4"] = "value4"
	b["abc2"]["key5"] = "value5"
	b["abc2"]["key6"] = "value6"
	fmt.Println(b)
}
func main() {
	testMap()
	testMapTwo()
}
