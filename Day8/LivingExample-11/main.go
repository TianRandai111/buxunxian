package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	exitChan := make(chan bool, 2)
	go sendData(ch, exitChan)
	go getData(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
}

func sendData(ch chan string, e chan bool) {
	ch <- "Washing"
	ch <- "Tripol"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
	e <- true
}

func getData(ch chan string, e chan bool) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	e <- true
}
