package helper

import "fmt"

func Total(price int){
	pajak := (price / 100) * 10
	normalPrice := price - pajak
	fmt.Println("harga normal :", normalPrice )
	fmt.Println("Pajak PPN :", pajak)
	fmt.Println("Total PPN :", price)
}

