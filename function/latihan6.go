package function

import "strconv"

func TotalPpn(belanja map[string]int) string {
	result := 0
	for _, value := range belanja {
		var price = (100 * value) / 110

		var ppn = value - price
		result += ppn
	}
	return "total PPN yang diterima sebesar " + strconv.Itoa(result)
}
