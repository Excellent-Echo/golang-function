package main

import (
	"fmt"
	"strconv"
)

//task 1
func Total(total float32) {
	var ppn, harga float32
	ppn = (total * 0.1)
	harga = total - ppn
	fmt.Println(ppn, harga)
}

//task 2
func changeToIdr(nominal int) {
	nom := strconv.Itoa(nominal)
	//curency := "IDR"
	leng := len(nom)
	var res string
	for i := 0; i <= leng-1; i++ {

		if i%3 == 0 {
			res += string(nom[i]) + ","
		} else {
			res += string(nom[i])
		}
	}
	fmt.Println("IDR " + res + "00")

}

func GenapGanjil(numbers ...int) {
	var even, odd int
	for _, number := range numbers {
		if number%2 == 0 {
			even++
		} else if number%2 != 0 {
			odd++
		}
	}
	// fmt.Println(odd)
	// fmt.Println(even)
	if even > odd {
		fmt.Println("angka terbanyak adalah genap")
	} else if even < odd {
		fmt.Println("angka terbanyak adalah ganjil")
	} else {
		fmt.Println("angka ganjil dan genap sama banyak")
	}
}

//task 4
func donationChecker(data map[string]string) {
	var err error

	nama := data["name"]
	umur, err := strconv.Atoi(data["umur"])
	jarak, err := strconv.Atoi(data["jarakRumah"])
	berkeluarga := data["berkeluarga"]

	if err == nil {
		if umur >= 20 && jarak <= 100 && berkeluarga == "ya" {
			fmt.Println(nama + " layak mendapat bantuan dari pemerintah")
		} else {
			fmt.Println(nama + " tidak layak mendapat bantuan dari pemerintah")
		}
	} else {
		fmt.Println(err.Error())
	}
}

//task 6
func countPPN(datas map[string]int) {
	var sum int
	for _, data := range datas {

		sum += data
	}
	res := sum / 11
	fmt.Printf("Total PPN yang diteruma sebesar %d ", res)
}

func main() {
	var data = map[string]string{
		"name":        "andi",
		"umur":        "30",
		"jarakRumah":  "50",
		"berkeluarga": "ya",
	}

	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}

	//task 1
	Total(110000)

	//task 2
	changeToIdr(1000000)

	//task 3
	GenapGanjil(1, 2, 3, 4, 5)

	//task 4
	donationChecker(data)

	// task 5
	// Skip

	//task 6
	countPPN(belanja1)
}
