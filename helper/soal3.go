package helper

func GenapGanjil(numbers ...int) (res string) {
	if len(numbers) < 1 {
		res = "tidak ada angka"
		return
	}
	var genap, ganjil int
	for _, number := range numbers {
		if number%2 == 0 {
			genap++
		} else {
			ganjil++
		}
	}
	if genap > ganjil {
		res = "Angka terbanyak genap"
	} else if ganjil > genap {
		res = "Angka terbanyak ganjil"
	} else {
		res = "Angka genap dan ganjil sama banyak"
	}
	return
}
