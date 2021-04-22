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
	nama := data["name"]
	umur, x := strconv.Atoi(data["umur"])
	jarak, y := strconv.Atoi(data["jarakRumah"])
	berkeluarga := data["berkeluarga"]

	if umur >= 20 && jarak <= 100 && berkeluarga == "ya" {
		fmt.Println(nama + " layak mendapat bantuan dari pemerintah")
	} else {
		fmt.Println(nama + " tidak layak mendapat bantuan dari pemerintah")
	}
	fmt.Println(x, y)
}

func main() {
	var data = map[string]string{
		"name":        "andi",
		"umur":        "30",
		"jarakRumah":  "50",
		"berkeluarga": "ya",
	}

	//task 1
	Total(110000)

	//task 2
	changeToIdr(1000000)

	//task 3
	GenapGanjil(1, 2, 3, 4, 5)

	//task 4
	donationChecker(data)
}
