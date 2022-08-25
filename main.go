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
	switch opt {
	case "a":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price: ", reader)
		fmt.Println(name, price)
	case "s":
		fmt.Println("saved")
	case "t":
		tip, _ := getInput("tip amount: ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("invalid")
		promptOptions(b)
	}

}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill.format())
}
