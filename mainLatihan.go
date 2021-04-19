package main

import (
	"fmt"
	"functionGo/latihan"
)

func main() {
	fmt.Println("HItung Pajak")
	latihan.HitungPajak(100000)
	fmt.Println("=================")

	latihan.CekGenapGanjil(1, 3, 6, 6, 2, 6)

}
