package latihan

import (
	"errors"
	"strconv"
)

func ChangeToIdr(nominal string) (string, error) {
	input, err := strconv.Atoi(nominal)

	idr := strconv.Itoa(input)
	thousand := 3
	if input < 0 {
		thousand++
	}
	for i := len(idr); i > thousand; {
		i -= 3
		idr = idr[:i] + "," + idr[i:]
	}

	if err != nil {
		return "", errors.New("Input incorrect: Please input a number")
	} else {
		return "IDR " + idr + ",00", nil
	}
}
