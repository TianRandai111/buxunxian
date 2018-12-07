package main

func main() {
	var a Link
	for i := 0; i < 10; i++ {
		//a.InsertHead(i)
		a.InsertTail(i)
	}
	a.Trans()
}
