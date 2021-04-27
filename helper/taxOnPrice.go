package helper

import "fmt"

func FindPPN(total int) {
	beforePpn := total * 10 / 11
	ppn := total / 11
	fmt.Println("Sebelum Pajak :", beforePpn)
	fmt.Println("Pajak : ", ppn)
}