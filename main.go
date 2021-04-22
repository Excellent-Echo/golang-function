package main

import (
	"fmt"
	"strconv"
)

//task 1
func Total(total float32) {
	var ppn, harga float32
	ppn = (total * 0.1)
	harga = total - ppn
	fmt.Println(ppn, harga)
}

//task 2
func changeToIdr(nominal int) {
	nom := strconv.Itoa(nominal)
	//curency := "IDR"
	leng := len(nom)
	var res string
	for i := 0; i <= leng-1; i++ {

		if i%3 == 0 {
			res += string(nom[i]) + ","
		} else {
			res += string(nom[i])
		}
	}
	fmt.Println("IDR " + res + "00")

}

func main() {
	//task 1
	Total(110000)

	//task 2
	changeToIdr(1000000)
}
