package main

import (
	"fmt"
)

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type Hello interface {
	Hello()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", (*p).Name)
}
func (p *BMW) Hello() {
	fmt.Printf("%s is hello\n", (*p).Name)
}

type BYD struct {
	Name string
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BYD) DiDi() {
	fmt.Printf("%s is didi\n", (*p).Name)
}

func main() {
	var (
		cat  Carer
		bmw  BMW
		byd  BYD
		test Hello
	)

	bmw.Name = "BMW"
	cat = &bmw
	cat.Run()
	cat.DiDi()
	test = &bmw
	test.Hello()
	byd.Name = "BYD"
	cat = &bmw
	cat.Run()
	cat.DiDi()
}
