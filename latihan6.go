package main

import "fmt"

func main() {
	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}
		// sepatu := ["sepatu"]
		// jaket := ["jaket"]
		// topi := ["topi"]
		// celana := ["celana"]
		// sweater := ["sweater"]
		// kausKaki := ["kausKaki"]
		// sabuk := ["sabuk"]

		var totalbersih = sepatu + jaket + topi + celana + sweater + kausKaki + sabuk
		var pajak = totalbersih * 10 / 100
		total := totalbersih + pajak

		fmt.Printf("total PPN yang diterima sebesar : %v ", total)
	}
}