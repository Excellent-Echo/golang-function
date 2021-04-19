package helper

func Total(bayarTotal int) (int, int) {
	bayar := bayarTotal * 10 / 11
	ppn := bayarTotal / 11

	return ppn, bayar
}
