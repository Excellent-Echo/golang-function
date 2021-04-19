package latihan

func Belanja(total int) (ppn int, beforePpn int) {
	beforePpn = total * 10 / 11
	ppn = total / 11
	return
}
