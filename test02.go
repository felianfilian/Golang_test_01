package main

import "fmt"

func sayHello(n string) {
	fmt.Printf("Hello %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	cycleNames([]string{"mario", "peach", "bowser"}, sayHello)
}
