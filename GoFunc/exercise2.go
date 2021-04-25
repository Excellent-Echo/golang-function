package GoFunc

import (
	"errors"
	"strconv"
)

func ChangeToIdr(input string) (string, error) {
	nominal, err := strconv.Atoi(input)

	switch {
	case err != nil:
		return "", errors.New("Incorrect input")
	}

	money := strconv.Itoa(nominal)
	thousand := 3

	switch {
	case nominal < 0:
		thousand++
	}

	for i := len(money); i > thousand; {
		i -= 3
		money = money[:i] + "." + money[i:]
	}

	return "IDR " + money + ",00", nil
}