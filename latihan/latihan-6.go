package latihan

import (
	"strconv"
)

func Ppn(data map[string]int) string {
	sum := 0

	for _, val := range data {
		ppn := val / 11
		sum += ppn
	}
	return strconv.Itoa(sum)
}
