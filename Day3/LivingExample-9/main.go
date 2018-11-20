package main

import (
	"fmt"
	"time"
)

func Now_Time() {
	D_T := time.Now()
	fmt.Println(D_T.Format("2006/1/02 15:04:05"))
}

func main() {
	Now_Time()
}
