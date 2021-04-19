package helper

import "strconv"

func ChangeToIdr(nominal int) string {
	idr := strconv.Itoa(nominal)
	thousand := 3
	if nominal < 0 {
		thousand++
	}
	for i := len(idr); i > thousand; {
		i -= 3
		idr = idr[:i] + "," + idr[i:]
	}
	return "IDR " + idr + ",00"
}
