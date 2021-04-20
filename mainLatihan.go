package main

import (
	"fmt"
	"functionGo/latihan"
)

func main() {
	fmt.Println("Hitung Pajak")
	latihan.HitungPajak(100000)
	fmt.Println("=================")
	fmt.Println("Genap Ganjil")
	latihan.CekGenapGanjil(1, 3, 6, 6, 2, 6)
	fmt.Println("================")
	fmt.Println("Change To IDR")
	latihan.ChangeRupiah(100000)

}
