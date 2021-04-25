package main

import (
	"fmt"
	"functionGo/latihan"
)

func main() {

	var inputPajak string
	fmt.Println("Hitung Pajak")
	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scan(&inputPajak)
	resultTax, resultSum, err := latihan.HitungPajak(inputPajak)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("PPN: %v, Total Belanja: %v\n", resultTax, resultSum)
	}
	fmt.Println("=================")
	fmt.Println("Genap Ganjil")
	latihan.CekGenapGanjil(1, 3, 6, 6, 2, 6)
	fmt.Println("================")
	fmt.Println("Change To IDR")
	latihan.ChangeRupiah(100000)

}
