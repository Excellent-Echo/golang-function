package function

func TotalPrice(total int) (int, int) {
	// var tax = int(10 / 100)
	var price = (100 * total) / 110

	var ppn = total - price
	return ppn, price

}
