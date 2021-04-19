package helper

func TotalPPN(data map[string]int) (ppn int) {
	for _, total := range data {
		ppn += total / 11
	}
	return
}
