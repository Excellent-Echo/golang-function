package main

import "fmt"

func GenapGanjil(num ...int) string {
	var (
		status        string
		ganjil, genap int
	)

	if num == nil {
		status = "tidak ada angka"
	}

	for _, val := range num {
		if val%2 != 0 {
			ganjil++
		} else if val%2 == 0 {
			genap++
		} else {
			status = "hasil tidak valid"
		}
	}

	if ganjil > genap {
		status = "angka terbanyak adalah ganjil"
	} else if genap > ganjil {
		status = "angka terbanyak genap"
	}

	return status
}

func main() {
	fmt.Println(GenapGanjil(1, 2, 3, 4, 5))
	fmt.Println(GenapGanjil(4, 2))
	fmt.Println(GenapGanjil(10, 20, 30, 13))
	fmt.Println(GenapGanjil(30, 13, 13, 77, 33, 55, 17, 13))
	fmt.Println(GenapGanjil())
}
