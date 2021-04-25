package helper

import "fmt"

func GenapGanjil(number ...int) {
	if number == nil {
		fmt.Println("Tidak ada angka")
	}

	var (
		genap, ganjil int
	)

	for _, value := range number {
		if value%2 == 0 {
			genap++
			//fmt.Println("ini adalah genap", value)
		}
		if value%2 == 1 {
			ganjil++
			//fmt.Println("ini adalah ganjil", value)
		}
	}
	//fmt.Println(genap)
	//fmt.Println(ganjil)
	if genap > ganjil {
		fmt.Println("angka terbanyak adalah genap")
	}
	if ganjil > genap {
		fmt.Println("angka terbanyak adalah ganjil")
	}
}
