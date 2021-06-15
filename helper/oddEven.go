package helper

import "fmt"

func OddEven(numbers ...int) {
	even, odd := 0, 0

	for _, num := range numbers {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		fmt.Println("Angka terbanyak adalah genap")
	} else if even < odd {
		fmt.Println("Angka terbanyak adalah ganjil")
	} else {
		fmt.Println("Unknown")
	}

}