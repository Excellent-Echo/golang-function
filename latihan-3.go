package main

import "fmt"

func GenapGanjil(number ...int) {

	if number == nil {
		fmt.Println("Tidak ada angka")
	}

	var (
		genap, ganjil int
	)

	for _, value := range number {
		if value % 2 == 0 {
			genap++
			//fmt.Println("ini adalah genap", value)
		}

		if value % 2 == 1 {
			ganjil++
			//fmt.Println("ini adalah ganjil", value)
		}
	}
	//fmt.Println(genap)
	//fmt.Println(ganjil)

	if genap > ganjil {
		fmt.Println("Genap lebih banyak")
	}

	if ganjil > genap {
		fmt.Println("Ganjil lebih banyak")
	}
}

func main(){
	GenapGanjil(1,2,3,4,5)
	GenapGanjil(4,2)
	GenapGanjil(10,20,30,13)
	GenapGanjil(30,13,13,77,33,55,17,13)
	GenapGanjil()
}