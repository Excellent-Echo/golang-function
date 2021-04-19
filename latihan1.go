package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("masukkan input : ")
	scanner.Scan()
	hargaawal, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	var pajak = hargaawal * 10 / 100
	total := hargaawal + pajak
	fmt.Printf("Harga awal : %v, Harga PPN : %v, Harga Total Harga : %v", hargaawal, pajak, total)

}
