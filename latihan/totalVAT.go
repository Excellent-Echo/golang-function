package latihan

func TotalVAT(data map[string]int) int {
	total := 0

	for _, v := range data {
		total += v
	}
	return total
}
