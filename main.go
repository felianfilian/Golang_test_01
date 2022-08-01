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
		fmt.Println("item added")
	case "s":
		fmt.Println("saved")
	case "t":
		fmt.Println("tip added")
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
