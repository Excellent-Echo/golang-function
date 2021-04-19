package main

import (
	"fmt"
)

func GenapGanjil(numbers ...int) string {
	countGanjil := 0
	countGenap := 0

	if len(numbers) <= 0 {
		return "tidak ada angka"
	}

	for _, num := range numbers {
		if num%2 == 0 {
			countGenap++
		} else {
			countGanjil++
		}
	}

	if countGanjil > countGenap {
		return "angka terbanyak adalah ganjil"
	} else if countGenap > countGanjil {
		return "angka terbanyak adalah genap"
	}

	return ""
}

func main() {
	fmt.Println(GenapGanjil(1, 2, 3, 4, 5))
	fmt.Println(GenapGanjil(4, 2))
	fmt.Println(GenapGanjil(10, 20, 30, 13))
	fmt.Println(GenapGanjil(30, 13, 13, 77, 33, 55, 17, 13))
	fmt.Println(GenapGanjil())
}
