package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//Input data
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("masukkan input : ")
	scanner.Scan()
	hargaawal, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	//Hitung
	pajak := hargaawal * 10 / 100
	total := hargaawal + pajak

	fmt.Printf("Harga awal : Rp.%v, Harga PPN : Rp.%v, Total Harga+ppn : Rp.%v", hargaawal, pajak, total)

}
