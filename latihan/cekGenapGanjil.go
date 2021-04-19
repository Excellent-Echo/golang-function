package latihan

import "fmt"

func CekGenapGanjil(numbers ...int) {
	checkEven := 0
	checkOdd := 0

	for _, num := range numbers {
		if num%2 != 0 {
			checkEven++
		} else {
			checkOdd++
		}
	}

	if checkEven > checkOdd {
		fmt.Println("angka terbanyak adalah ganjil")
	} else if checkOdd > checkEven {
		fmt.Println("angka terbanyak adalah genap")
	} else {
		fmt.Println("Tidak ada angka")
	}

}
