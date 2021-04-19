package latihan

func ValueAddedTax(total int) (int, int) {
	x := (float64(100) / 110)

	price := float64(total) * x
	vat := total - int(price)

	return vat, int(price)
}
