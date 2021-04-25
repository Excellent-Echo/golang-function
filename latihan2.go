package main

import (
	"fmt"
	"strconv"
)

func main() {

	changeToIdr(1000000)

}
func changeToIdr(number int) {

	numstr := strconv.Itoa(number) //hasil inputan diconvert dahulu menjadi string, agar bisa ditambahkan ","
	var tampung string             // deklarasi variabel untuk menampung hasil akhir yang telah dilooping ditambahkan ","

	for i := 0; i <= len(numstr)-1; i++ {
		// perulangan jika saat nilai i/3=0 maka ditambahkan ","
		if i%3 == 0 {
			tampung += string(numstr[i]) + ","
		} else { // jika i/3 != 0 maka akan nilainya ditampung biasa tanpa ","
			tampung += string(numstr[i])
		}
		fmt.Println("IDR " + tampung + "00")
	}
}
