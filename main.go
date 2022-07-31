package main

import "fmt"

func main() {
	myBill := newBill("Luigi")
	fmt.Println(myBill.format())
}
