package main

import (
	"fmt"
	"functionGo/helper"
)

func main() {

	//Latihan 1

	helper.TaxOnPrice(110000)

	fmt.Println("\n\n#######################################")

	

	//Latihan 3

	num := helper.OddEven(1, 2, 3, 4, 5)
	fmt.Printf("%v", num)

	fmt.Println("\n\n#######################################")

}