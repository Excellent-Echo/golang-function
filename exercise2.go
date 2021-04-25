package main

import (
	"fmt"
	"functionGo/GoFunc"
)

func main() {
	var input string
	fmt.Print("Input number: ")
	fmt.Scanln(&input)
	result2, err := GoFunc.ChangeToIdr(input)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Println(result2)
	}
}
