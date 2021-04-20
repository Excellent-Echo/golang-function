package helper

func OddEven(numbers ...int) string {
	even, odd := 0, 0

	for _, num := range numbers {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		return "angka terbanyak adalah genap"
	} else if even < odd {
		return "angka terbanyak adalah ganjil"
	} else {
		return "tidak ada angka"
	}

}