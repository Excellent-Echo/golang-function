package helper

import (
	"fmt"
	"strconv"
)

func TotalPajak(belanja1 map[string]int)  {

	total := 0
	for _, value := range belanja1 {
		//fmt.Println(value)
		pajak := (value / 100) * 10
		//fmt.Println(pajak)
		total += pajak
	}

	toString := strconv.Itoa(total)

	fmt.Println("total PPN yang diterima sebesar " + toString)
}