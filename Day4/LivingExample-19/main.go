package main

import "fmt"

func test() {
	var a map[string]string = map[string]string{"hello": "world"}
	a = make(map[string]string, 10)
	a["hello"] = "world"  //插入和更新
	val, ok := a["hello"] //查找
	if ok {
		fmt.Printf("This is value:%s\n", val)
	} else {
		fmt.Printf("This is vulue not in map\n")
	}

	for k, v := range a { //遍历
		fmt.Println(k, v)
	}

	b := len(a) //长度
	fmt.Printf("map.a的长度：%d\n", b)

	delete(a, "hello") //删除
	b = len(a)         //长度
	fmt.Printf("delete.map.a的长度：%d", b)

}

func testMap(a map[string]map[string]string) {
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["password"] = "123456"
	a["zhangsan"]["nickname"] = "pangpang"

	for k, v := range a {
		fmt.Println(k, ":")
		for c, d := range v {
			fmt.Println("\t", c, ":", d)
		}
	}
	b := len(a) //长度
	fmt.Printf("delete.map.a的长度：%d\n", b)
	delete(a, "zhangsan")
	// value, ok := a["zhangsan"]
	// if ok {
	// 	value["password"] = "123456"
	// 	value["nickname"] = "pangpang"
	// } else {
	// 	a["zhangsan"] = make(map[string]string)
	// 	a["zhangsan"]["password"] = "123456"
	// 	a["zhangsan"]["nickname"] = "pangpang"
	// }
}
func main() {
	test()
	fmt.Printf("\n-------------------------------------------------------------------\n")
	var info map[string]map[string]string
	info = make(map[string]map[string]string, 10)
	info["zhangsan"] = make(map[string]string)
	info["zhangsan"]["password"] = "987654"
	info["zhangsan"]["nickname"] = "default"
	info["tom"] = make(map[string]string)
	info["tom"]["password"] = "871365"
	info["tom"]["nickname"] = "nihao"
	testMap(info)
	fmt.Println(info)
}
