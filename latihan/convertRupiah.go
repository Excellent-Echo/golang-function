package latihan

import (
	"errors"
	"strconv"
)

func ChangeRupiah(currency string) (string, error) {
	nominal, err := strconv.Atoi(currency)
	idr := strconv.Itoa(nominal)
	thousands := 3
	length := len(idr)

	if nominal < 0 {
		thousands++
	}

	for i := length; i > thousands; {
		i = i - 3
		idr = idr[:i] + "," + idr[i:]
	}

	if err != nil {
		return "", errors.New("Not a Number")
	} else {
		return "IDR " + idr + ",00", nil
	}

}
