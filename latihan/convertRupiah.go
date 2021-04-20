package latihan

import (
	"fmt"
	"strconv"
)

func ChangeRupiah(currency int) {
	idr := strconv.Itoa(currency)
	thousands := 3
	length := len(idr)
	if currency < 0 {
		thousands++
	}
	for i := length; i > thousands; {
		i = i - 3
		idr = idr[:i] + "," + idr[i:]
	}
	fmt.Println("IDR " + idr + ",00")
}
