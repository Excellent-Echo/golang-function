package helper

import (
	"errors"
	"fmt"
)

func validate(price int) (bool, error) {
	if price == 0 {
		return false, errors.New("Tidak boleh tanpa value")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Keterangan error :", r)
	} else {
		fmt.Println("Aplikasi berjalan sempurna")
	}
}

func Total(price int) {
	defer catch()

	if valid, err := validate(price); valid {
		pajak := price / 11
		normalPrice := price - pajak
		fmt.Println("harga normal :", normalPrice)
		fmt.Println("Pajak PPN :", pajak)
		fmt.Println("Total PPN :", price)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
