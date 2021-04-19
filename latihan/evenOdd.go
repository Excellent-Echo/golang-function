package latihan

func EvenOdd(numbers ...int) string {
	countEven, countOdd := 0, 0

	for _, num := range numbers {
		if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}

	if countEven > countOdd {
		return "angka terbanyak adalah genap"
	} else if countEven < countOdd {
		return "angka terbanyak adalah ganjil"
	} else {
		return "tidak ada angka"
	}

}
