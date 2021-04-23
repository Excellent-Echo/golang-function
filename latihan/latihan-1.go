package latihan

import (
	"errors"
	"strconv"
)

func Belanja(total string) (ppn int, beforePpn int, err error) {
	input, err := strconv.Atoi(total)

	if err != nil {
		return 0, 0, errors.New("Not a number")
	} else {
		beforePpn = input * 10 / 11
		ppn = input / 11
		return ppn, beforePpn, nil
	}
}
