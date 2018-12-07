package main

import (
	"fmt"
)

type Cart struct {
	weight int
	name   string
}

func (p *Cart) Run() {
	fmt.Println("Running")
}

//匿名字段
type Bike struct {
	Cart
	lunzi int
}

type Train struct {
	Cart
}

//实现string()
func (p *Train) String() string {
	str := fmt.Sprintf("[name]=[%s] weight=[%d]", p.name, p.weight)
	return str
}

//多重集成
type ATrain struct {
	c Cart
}

func main() {
	var a Bike
	a.Cart.weight = 100
	a.Cart.name = "奥迪"
	a.lunzi = 3

	fmt.Println(a)
	a.Run()

	var t Train
	t.Cart.weight = 100
	t.Cart.name = "奥迪"

	fmt.Println(t)
	t.Run()
	d := t.String()
	fmt.Println(d)

	var at ATrain
	at.c.weight = 1000
	at.c.name = "毛爷爷"

	fmt.Println(at)
	t.Run()

}
