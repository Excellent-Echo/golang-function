package main

import "fmt"

func changeToIdr(price int) string {
	fmt.Println("price", price)
	if price < 0 {
		return "-" + changeToIdr(-price)
	}

	if price < 1000 {
		return fmt.Sprintf("%d",price)
	}
	return changeToIdr(price / 1000) + "," + fmt.Sprintf("%03d",price % 1000)
}

func main() {
	fmt.Println(changeToIdr(1200000))
}