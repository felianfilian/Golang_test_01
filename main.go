package main

import "fmt"

func main() {
	myBill := newBill("Luigi")
	myBill.addItem("cake", 3.99)
	myBill.addItem("coffee", 1.99)
	myBill.updateTip(2)
	fmt.Println(myBill.format())
}
