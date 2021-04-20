package main

import (
	"fmt"
	"strconv"
)

func ChangeToIdr(money int) string {
	idr := strconv.Itoa(money)

	nol := 3

	if money < 0 {
		nol++
	}

	for i := len(idr); i > nol; {
		i -= 3
		idr = idr[:i] + "." + idr[i:]
	}

	return idr
}

func main() {
	fmt.Println("IDR", ChangeToIdr(50000), ",00")
}
