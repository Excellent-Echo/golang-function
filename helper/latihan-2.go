package helper

import "fmt"

func ChangeToIdr(price int) string {
	if price < 0 {
		return "-" + ChangeToIdr(-price)
	}

	if price < 1000 {
		return fmt.Sprintf("%d", price)
	}

	result := ChangeToIdr(price/1000) + "," + fmt.Sprintf("%03d", price%1000)

	final := "IDR " + result + "," + "00"
	fmt.Println("Hasil", final)
	return result
}
