package main

import "fmt"

func Total(price int){
	pajak := (price / 100) * 10
	normalPrice := price - pajak

	fmt.Println("harga normal :", normalPrice )
	fmt.Println("Pajak PPN :", pajak)
	fmt.Println("Total PPN :", price)
}

func main(){
	Total(110000)
	Total(220000)
	Total(500000)
	Total(1000000)
}

