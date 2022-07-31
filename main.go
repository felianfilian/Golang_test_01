package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create new bill", reader)
	b := newBill(name)
	fmt.Println("Created the bill: ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("ChooseOption:\na - add item\ns - save bill\nt - add tip", reader)
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill.format())
	// myBill := newBill("Luigi")
	// myBill.addItem("cake", 3.99)
	// myBill.addItem("coffee", 1.99)
	// myBill.updateTip(2)
	// fmt.Println(myBill.format())
}
