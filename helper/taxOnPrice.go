package helper

import "fmt"

func TaxOnPrice(price int){
	tax := (price * 10) / 100
	priceBeforeTax := price - tax
	fmt.Printf("Harga normal : %v dan Pajak : %v", priceBeforeTax, tax)
}