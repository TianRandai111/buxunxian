package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

type Train struct {
	Car
	Start time.Time
	int
}

func main() {
	var t Train
	/* 	t.Name = "train"
	   	t.Age = 100 */
	t.Car.Name = "train"
	t.Car.Age = 100
	t.int = 200
	fmt.Println(t)
}
