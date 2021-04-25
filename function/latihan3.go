package function

func GenapGanjil(number ...int) string {
	ganjil := 0
	genap := 0

	for _, value := range number {
		if value%2 == 0 {
			genap++
		} else {
			ganjil++
		}
	}
	if len(number) == 0 {
		return "tidak ada angka"
	} else if genap < ganjil {
		return "angka terbanyak adalah ganjil"
	} else {
		return "angka terbanyak adalah genap"
	}
}
