package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var mytemplate *template.Template

func simpleServer(w http.ResponseWriter, request *http.Request) {
	fmt.Println("hello hello")
	//fmt.Fprintf(w, "hello")
	var arr []Person
	p := Person{Name: "Mary", Age: 18}
	p1 := Person{Name: "Mary1", Age: 181}
	p2 := Person{Name: "Mary2", Age: 182}
	arr = append(arr, p)
	arr = append(arr, p1)
	arr = append(arr, p2)
	mytemplate.Execute(w, arr)

}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func initTemplat(filename string) (err error) {
	mytemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	//初始化模板
	initTemplat("./index.html")
	http.HandleFunc("/user/info", simpleServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Println("http listen failed")
	}
}
