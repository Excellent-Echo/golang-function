package main

import "fmt"

func Total(subtotal int) (int, int) {
	ppn := (subtotal * 10) / 100
	total := subtotal - ppn
	return ppn, total
}

func main() {
	fmt.Println(Total(50000))
}
