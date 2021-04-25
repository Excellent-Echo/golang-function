package main

import (
	"fmt"
	"functionGo/GoFunc"
)

func main() {
	var input string
	fmt.Print("Input nominal: ")
	fmt.Scanln(&input)
	result, err := GoFunc.ChangeToIdr(input)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	default:
		fmt.Println(result)
	}
}
