package main

import "fmt"

func Total(total int) (int, int) {
	ppn := (total * 10) / 100
	subtotal := total - ppn
	return ppn, subtotal
}

func main() {
	fmt.Println(Total(50000))
}
