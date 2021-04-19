package main

import "fmt"

func belanja(total int) (ppn int, beforePpn int) {
	beforePpn = total * 10 / 11
	ppn = total / 11
	return
}

func main() {
	// fmt.Println(belanja(110000))
	// fmt.Println(belanja(220000))
	// fmt.Println(belanja(500000))
	fmt.Println(belanja(1000000))
}
