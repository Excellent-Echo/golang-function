package main

import (
	"fmt"
	"functionGo/helper"
)


func main() {

	//Latihan 1 : PPN
	helper.FindPPN(500000)
	fmt.Println("==================================")

	/*Latihan 2 : Change to IDR
	helper.ChangeToIdr("25000")
	fmt.Println("==================================")*/ //not finished

	//Latihan 3 : oddEven
	helper.OddEven(2, 3, 4, 1, 5)
	fmt.Println("==================================")

	


}