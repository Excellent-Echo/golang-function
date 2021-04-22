package main

import "fmt"

//task 1
func Total(total float32) {
	var ppn, harga float32
	ppn = (total * 0.1)
	harga = total - ppn
	fmt.Println(ppn, harga)
}

func main() {
	//task 1
	Total(110000)
}
