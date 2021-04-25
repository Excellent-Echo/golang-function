package GoFunc

import (
	"errors"
	"strconv"
)

func Total(input string) (tax int, price int, err error) {
	priceTotal, err := strconv.Atoi(input)

	switch {
	case err != nil:
		return 0, 0, errors.New("Incorrect input")
	}

	tax = priceTotal / 11
	price = priceTotal - tax

	return tax, price, nil
}