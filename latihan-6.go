package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var belanja1 = map[string]int{
		"sepatu" : 1100000,
		"jaket" : 2200000,
		"topi" : 594000,
		"celana" : 803000,
		"sweater" : 330000,
		"kausKaki" : 110000,
		"sabuk" : 55000,
	}

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