package main

import (
	"fmt"
	"functionGo/GoFunc"
)

func main() {
	var input string
	fmt.Printf("Input total shopping: ")
	fmt.Scanln(&input)
	tax, price, err := GoFunc.Shopping(input)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Printf("Tax: %v, Price: %v\n", tax, price)
	}
}
