package main

import (
	"fmt"
	"functionGo/latihan"
)

func main() {
	fmt.Println("Latihan 1")
	// Latihan 1
	var input string
	fmt.Print("Input total belanja: ")
	fmt.Scanln(&input)
	pajak, harga, err := latihan.Belanja(input)
	// pajak, harga := latihan.Belanja(220000)
	// pajak, harga := latihan.Belanja(500000)
	// pajak, harga := latihan.Belanja(1000000)

	// Handle error
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Pajak: %v, Harga: %v\n", pajak, harga)
	}

	fmt.Println("-------------------------------")

	// fmt.Println("Latihan 2")
	// // Latihan 2
	// fmt.Println(latihan.ChangeToIdr(100000))
	// fmt.Println(latihan.ChangeToIdr(1200000))
	// fmt.Println(latihan.ChangeToIdr(140350000))

	// fmt.Println("-------------------------------")

	// fmt.Println("Latihan 3")
	// // Latihan 3
	// fmt.Println(latihan.GenapGanjil(1, 2, 3, 4, 5))
	// fmt.Println(latihan.GenapGanjil(4, 2))
	// fmt.Println(latihan.GenapGanjil(10, 20, 30, 13))
	// fmt.Println(latihan.GenapGanjil(30, 13, 13, 77, 33, 55, 17, 13))
	// fmt.Println(latihan.GenapGanjil())

	// fmt.Println("-------------------------------")

	// fmt.Println("Latihan 4")
	// // Latihan 4
	// var data = map[string]string{
	// 	"name":        "andi",
	// 	"umur":        "30",
	// 	"jarakRumah":  "50",
	// 	"berkeluarga": "ya",
	// }
	// var data2 = map[string]string{
	// 	"name":        "santi",
	// 	"umur":        "19",
	// 	"jarakRumah":  "80",
	// 	"berkeluarga": "ya",
	// }
	// var data3 = map[string]string{
	// 	"name":        "budi",
	// 	"umur":        "45",
	// 	"jarakRumah":  "120",
	// 	"berkeluarga": "ya",
	// }

	// latihan.Bantuan(data)
	// latihan.Bantuan(data2)
	// latihan.Bantuan(data3)

	// fmt.Println("-------------------------------")

	// fmt.Println("Latihan 5")
	// // Latihan 5
	// fmt.Println(latihan.ChangeNumtoStr(100000))
	// fmt.Println(latihan.ChangeNumtoStr(111000))
	// fmt.Println(latihan.ChangeNumtoStr(5124000))
	// fmt.Println(latihan.ChangeNumtoStr(1543))
	// fmt.Println(latihan.ChangeNumtoStr(1234678))
	// fmt.Println(latihan.ChangeNumtoStr(0))

	// fmt.Println("-------------------------------")

	// fmt.Println("Latihan 6")
	// // Latihan 6
	// var belanja1 = map[string]int{
	// 	"sepatu":   1100000,
	// 	"jaket":    2200000,
	// 	"topi":     594000,
	// 	"celana":   803000,
	// 	"sweater":  330000,
	// 	"kausKaki": 110000,
	// 	"sabuk":    55000,
	// }

	// totalPpn := latihan.Ppn(belanja1)
	// fmt.Println("total PPN yang diterima sebesar " + totalPpn)

}
