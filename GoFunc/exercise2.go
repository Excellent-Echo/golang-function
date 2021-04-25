package GoFunc

import (
	"errors"
	"strconv"
)

func ChangeToIdr(nominal string) (string, error) {
	input, err := strconv.Atoi(nominal)

	switch {
	case err != nil:
		return "", errors.New("Input incorrect: Please input a number")
	}

	idr := strconv.Itoa(input)
	thousand := 3

	switch {
	case input < 0:
		thousand++
	}
	for i := len(idr); i > thousand; {
		i -= 3
		idr = idr[:i] + "," + idr[i:]
	}

	return "IDR " + idr + ",00", nil
}