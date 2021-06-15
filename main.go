package main

import (
	"fmt"
	"functionGo/helper"
)


func main() {

	//Latihan 1 : PPN
	fmt.Println("============LATIHAN 1==========")
	helper.TaxOnPrice(500000)

	//Latihan 2 : Change to IDR
	fmt.Println("============LATIHAN 2==========")
	helper.ToIdr(25000)
	

	//Latihan 3 : oddEven
	fmt.Println("============LATIHAN 3==========")
	helper.OddEven(2, 3, 4, 1, 5)

	//Latihan 4 : Stimulus Checks
	var data = map[string]string{
		"name":        "andi",
		"umur":        "30",
		"jarakRumah":  "50",
		"berkeluarga": "ya",
	}
	var data2 = map[string]string{
		"name":        "santi",
		"umur":        "19",
		"jarakRumah":  "80",
		"berkeluarga": "ya",
	}
	var data3 = map[string]string{
		"name":        "budi",
		"umur":        "45",
		"jarakRumah":  "120",
		"berkeluarga": "ya",
	}

	//Latihan 4 : Stimulus Check
	fmt.Println("============LATIHAN 4==========")
	helper.StimulusChecks(data)
	helper.StimulusChecks(data2)
	helper.StimulusChecks(data3)

	//Latihan 5 : Int To String
	fmt.Println("============LATIHAN 5==========")
	helper.IntToString(2532000)

	//Latihan 6 : Total PPN
	fmt.Println("============LATIHAN 5==========")
	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}

	helper.TotalTax(belanja1)





}