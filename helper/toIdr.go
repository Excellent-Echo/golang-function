package helper

import (
	"fmt"
	"strconv"
)

func ToIdr(nominal int)  {
	idr := strconv.Itoa(nominal)
	thousand := 3
	if nominal < 0 {
		thousand++
	}
	for i := len(idr); i > thousand; {
		i -= 3
		idr = idr[:i] + "," + idr[i:]
	}

	fmt.Println("IDR " + idr + ",00") 
}