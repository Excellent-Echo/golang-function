package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	moneychanger := accounting.Accounting{Symbol: "IDR", Precision: 2}
	fmt.Println(moneychanger.FormatMoney(123456789.213123))
	fmt.Println(moneychanger.FormatMoney(12345678))

}
