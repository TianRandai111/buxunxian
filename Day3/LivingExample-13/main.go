package main

import (
	"fmt"
	"strconv"
)

func transfrom(s *string) {
	i, err := strconv.Atoi(*s)
	if err != nil {
		fmt.Println("con not sconvert to int")
	} else {
		Info_type := fmt.Sprintf("%T", i)
		fmt.Printf("转换完毕，Value:%d,Type:%s", i, Info_type)
	}
	return
}

func main() {
	var In_int string
	fmt.Scanf("%s", &In_int)
	transfrom(&In_int)
}
