package latihan

import "fmt"

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
		fmt.Println("angka terbanyak adalah ganjil")
	} else if countGenap > countGanjil {
		fmt.Println("angka terbanyak adalah genap")
	}

	return ""
}
