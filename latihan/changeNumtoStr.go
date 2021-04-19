package latihan

import (
	"fmt"
	"strings"
)

func ChangeNumtoStr(num int) string {
	var str string

	unit := [12]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

	if num < 12 {
		str = unit[num]
	} else if num < 20 {
		str = fmt.Sprintf("%s belas", ChangeNumtoStr(num-10))
	} else if num < 100 {
		str = fmt.Sprintf("%s puluh %s", ChangeNumtoStr(num/10), ChangeNumtoStr(num%10))
	} else if num < 200 { // ratus
		str = fmt.Sprintf("seratus %s", ChangeNumtoStr(num-100))
	} else if num < 1000 {
		str = fmt.Sprintf("%s ratus %s", ChangeNumtoStr(num/100), ChangeNumtoStr(num%100))
	} else if num < 2000 { // ribu
		str = fmt.Sprintf("seribu %s", ChangeNumtoStr(num-1000))
	} else if num < 1000000 {
		str = fmt.Sprintf("%s ribu %s", ChangeNumtoStr(num/1000), ChangeNumtoStr(num%1000))
	} else if num < 2000000 { // juta
		str = fmt.Sprintf("satu juta %s", ChangeNumtoStr(num-1000000))
	} else if num < 1000000000 {
		str = fmt.Sprintf("%s juta %s", ChangeNumtoStr(num/1000000), ChangeNumtoStr(num%1000000))
	} else if num < 2000000000 { // milyar
		str = fmt.Sprintf("satu milyar %s", ChangeNumtoStr(num-1000000000))
	} else if num < 1000000000000 {
		str = fmt.Sprintf("%s milyar %s", ChangeNumtoStr(num/1000000000), ChangeNumtoStr(num%1000000000))
	} else if num < 2000000000000 { // triliun
		str = fmt.Sprintf("satu triliun %s", ChangeNumtoStr(num-1000000000000))
	} else if num < 1000000000000000 {
		str = fmt.Sprintf("%s triliun %s", ChangeNumtoStr(num/1000000000000), ChangeNumtoStr(num%1000000000000))
	}
	return strings.TrimSpace(str)

}
