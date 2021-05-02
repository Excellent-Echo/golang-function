package helper

import (
	"fmt"
	"strconv"
)

func TotalTax(data map[string]int) {
	sum := 0

	for _, val := range data {
		tax := val / 11
		sum += tax
	}
	fmt.Println(strconv.Itoa(sum))
}