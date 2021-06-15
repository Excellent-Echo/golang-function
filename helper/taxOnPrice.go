package helper

import "fmt"

func TaxOnPrice(total int) {
	beforeTax := total * 10 / 11
	ppn := total / 11
	fmt.Println("Sebelum Pajak :", beforeTax)
	fmt.Println("Pajak : ", ppn)
}