package main

import (
	"fmt"
	"strconv"
)

func changeToIdr(nominal int) string {
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

func main() {
	fmt.Println(changeToIdr(100000))
	fmt.Println(changeToIdr(1200000))
	fmt.Println(changeToIdr(140350000))
}
