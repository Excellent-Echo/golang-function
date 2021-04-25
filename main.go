package main

import (
	"fmt"
	"functionGo/function"
)

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

var belanja1 = map[string]int{
	"sepatu":   1100000,
	"jaket":    2200000,
	"topi":     594000,
	"celana":   803000,
	"sweater":  330000,
	"kausKaki": 110000,
	"sabuk":    55000,
}

func main() {

	fmt.Println("Latihan 1")
	fmt.Println(function.TotalPrice("sdads"))
	fmt.Println(function.TotalPrice("220000"))
	fmt.Println(function.TotalPrice("500000"))
	fmt.Println("\n")

	fmt.Println("Latihan 2")
	fmt.Println("belum selesai\n")

	fmt.Println("latihan 3")
	fmt.Println(function.GenapGanjil(1, 2, 3, 4, 5))
	fmt.Println(function.GenapGanjil(4, 2))
	fmt.Println(function.GenapGanjil(10, 20, 30, 13))
	fmt.Println(function.GenapGanjil(30, 13, 13, 77, 33, 55, 17, 13))
	fmt.Println(function.GenapGanjil(), "\n")

	fmt.Println("Latihan 4")
	fmt.Println(function.Check(data))
	fmt.Println(function.Check(data2))
	fmt.Println(function.Check(data3))

	fmt.Println("Latihan 5")
	fmt.Println("belum selesai\n")

	fmt.Println("Latihan 6")
	fmt.Println(function.TotalPpn(belanja1))

}
