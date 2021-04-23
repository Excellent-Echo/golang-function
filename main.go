package main

import (
	"fmt"
	"functionGo/latihan"
)

func main() {
	// Latihan 1
	fmt.Println("Latihan 1")
	var input string
	fmt.Print("Input total belanja: ")
	fmt.Scanln(&input)
	pajak, harga, err := latihan.Belanja(input)
	// pajak, harga, err := latihan.Belanja(220000)
	// pajak, harga, err := latihan.Belanja(500000)
	// pajak, harga, err := latihan.Belanja(1000000)

	// Handle error latihan 1
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Pajak: %v, Harga: %v\n", pajak, harga)
	}

	fmt.Println("-------------------------------")

	// Latihan 2
	fmt.Println("Latihan 2")
	var input2 string
	fmt.Print("Input number: ")
	fmt.Scanln(&input2)
	result2, err := latihan.ChangeToIdr(input2)

	// Handle error latihan 2
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result2)
	}

	fmt.Println("-------------------------------")

	// Latihan 3
	fmt.Println("Latihan 3")
	// var input3 string
	// fmt.Print("Input some numbers: ")
	// fmt.Scanln(&input3)
	// result3, err := latihan.ChangeToIdr(input2)

	// latihan.GenapGanjil(1, 2, 3, 4, 5)

	// fmt.Println("-------------------------------")

	// Latihan 4
	fmt.Println("Latihan 4")
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

	latihan.Bantuan(data)
	latihan.Bantuan(data2)
	latihan.Bantuan(data3)

	fmt.Println("-------------------------------")

	// Latihan 5
	fmt.Println("Latihan 5")
	var input5 string
	fmt.Print("Input number: ")
	fmt.Scanln(&input5)
	result5, err := latihan.ChangeNumtoStr(input5)

	// Handle error latihan 5
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result5)
	}

	fmt.Println("-------------------------------")

	// Latihan 6
	fmt.Println("Latihan 6")
	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}

	totalPpn := latihan.Ppn(belanja1)
	fmt.Println("total PPN yang diterima sebesar " + totalPpn)

}
