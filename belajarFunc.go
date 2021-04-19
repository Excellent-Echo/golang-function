package main

import (
	"fmt"
	"functionGo/helper"
)

func main() {
	//Soal 1
	fmt.Printf("===Soal 1===\n\n")
	ppn, bayar := helper.Total(1000000)
	fmt.Printf("ppn : %v\nbayar : %v\n\n", ppn, bayar)

	//Soal2
	fmt.Printf("===Soal 2===\n\n")
	idr := helper.ChangeToIdr(1000000)
	fmt.Printf("%v\n\n", idr)

	//Soal3
	fmt.Printf("===Soal 3===\n\n")
	genapGanjil := helper.GenapGanjil(2, 3, 4, 5, 6)
	fmt.Printf("%v\n\n", genapGanjil)

	//Soal4
	var data = map[string]string{
		"name":        "andi",
		"umur":        "30",
		"jarakRumah":  "50",
		"berkeluarga": "ya",
	}

	var data2 = map[string]string{
		"name":        "santi",
		"umur":        "19",
		"jarakRumah":  "80",
		"berkeluarga": "ya",
	}

	var data3 = map[string]string{
		"name":        "budi",
		"umur":        "45",
		"jarakRumah":  "120",
		"berkeluarga": "ya",
	}

	checkBantuan1 := helper.BantuanCheck(data)
	checkBantuan2 := helper.BantuanCheck(data2)
	checkBantuan3 := helper.BantuanCheck(data3)
	fmt.Printf("===Soal 4===\n\n")
	fmt.Printf("%v\n%v\n%v\n\n", checkBantuan1, checkBantuan2, checkBantuan3)

	//Soal 5
	fmt.Printf("===Soal 5===\n\n")
	stringRupiah := helper.ChangeNumtoStr(125689)
	fmt.Println(stringRupiah)
	fmt.Println("")

	//Soal 6
	fmt.Printf("===Soal 6===\n\n")
	var belanja1 = map[string]int{
		"sepatu":   1100000,
		"jaket":    2200000,
		"topi":     594000,
		"celana":   803000,
		"sweater":  330000,
		"kausKaki": 110000,
		"sabuk":    55000,
	}
	totalPPN := helper.TotalPPN(belanja1)

	fmt.Printf("Total PPN : %v\n\n", totalPPN)
}

// hello.Hello()

// sum, avg, check := hello.Sum(10, 11, 12)

// fmt.Println(sum, avg, check)

// closure function
// var getTotal = func(numbers ...int) int {
// 	var sum int

// 	for _, num := range numbers {
// 		sum += num
// 	}

// 	return sum
// }

// function as parameter
// result1 := sayHelloFilter("anjing", filterName)
// result2 := sayHelloFilter("afista", filterName)

// fmt.Println(result1)
// fmt.Println(result2)

// IIFE
// var check = func(name string, age int) bool {

// 	if name == "" || age == 0 {
// 		return true
// 	} else {
// 		return false
// 	}

// }("afista", 23)

// fmt.Println(check)
// func sayHelloFilter(name string, filter Filt) string { // menerima sebuah fungsi bukan membuat fungsi sendiri
// 	return "hello " + filter(name)
// }

// func filterName(name string) string {
// 	if name == "anjing" || name == "babi" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }
