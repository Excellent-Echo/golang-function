package latihan

import (
	"errors"
	"strconv"
)

func HitungPajak(result string) (resTax int, resSum int, err error) {
	input, err := strconv.Atoi(result)

	if err != nil {
		return 0, 0, errors.New("Not a number")
	} else {
		resTax := (input / 100) * 1
		resSum := (input - resTax)
		return resTax, resSum, nil
	}
}
