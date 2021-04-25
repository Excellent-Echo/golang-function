package function

import (
	"errors"
	"strconv"
)

func TotalPrice(total string) (int, int, error) {
	totalPrice, err := strconv.Atoi(total)
	if err != nil {
		return 0, 0, errors.New("invalid input, must be number")
	}
	var price = (100 * totalPrice) / 110

	ppn := totalPrice - price

	return ppn, price, err

}
