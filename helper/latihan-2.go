package helper

import "fmt"

func ChangeToIdr(price int) string {
	fmt.Println("price", price)
	if price < 0 {
		return "-" + ChangeToIdr(-price)
	}
	if price < 1000 {
		return fmt.Sprintf("%d",price)
	}
	return ChangeToIdr(price / 1000) + "," + fmt.Sprintf("%03d",price % 1000)
}
