package main

import (
	"fmt"
	"functionGo/latihan"
)

func main() {
	//LATIHAN 1 ======================================================================

	// vat, price := latihan.ValueAddedTax(110000) // Harga 100000 + PPN sebesar 10000
	vat, price := latihan.ValueAddedTax(50000) // Harga 45454 + PPN sebesar 4546

	fmt.Printf("Harga sebelum pajak %v + PPN sebesar %v\n", price, vat)
	fmt.Println("=============================================================")

	//LATIHAN 2 ======================================================================

	//LATIHAN 3 ======================================================================

	// numbers := latihan.EvenOdd() // tidak ada angka
	// numbers := latihan.EvenOdd(5, 4, 3, 2, 1) // angka terbanyak adalah ganjil
	numbers := latihan.EvenOdd(22, 33, 44, 55, 66) // angka terbanyak adalah genap
	fmt.Printf("%v\n", numbers)
	fmt.Println("=============================================================")

	// LATIHAN 4

	// var data = map[string]string{
	// 	"name":        "andi",
	// 	"umur":        "30",
	// 	"jarakRumah":  "50",
	// 	"berkeluarga": "ya",
	// }

	var data2 = map[string]string{
		"name":        "santi",
		"umur":        "19",
		"jarakRumah":  "80",
		"berkeluarga": "ya",
	}

	// var data3 = map[string]string{
	// 	"name":        "budi",
	// 	"umur":        "45",
	// 	"jarakRumah":  "120",
	// 	"berkeluarga": "ya",
	// }

	result := latihan.GovFunding(data2)
	fmt.Println(result)
	fmt.Println("=============================================================")

	// LATIHAN 5
	results := latihan.ChangeNumtoStr(1234567890)
	fmt.Println(results)

	fmt.Println("=============================================================")

	//LATIHAN 6
	var belanja = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}

	totalPrice := latihan.TotalVAT(belanja)
	totalVAT, _ := latihan.ValueAddedTax(totalPrice) // total PPN yang diterima sebesar 472000

	fmt.Printf("total PPN yang diterima sebesar %v\n", totalVAT)
}
