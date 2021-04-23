package latihan

import (
"errors"
"fmt"	
"strconv"
)


func HitungPajak() {
	var nilai float32
	fmt.Print("Masukkan Nominal")
	mt.Scan(&nilai)
	resTax := (nilai / 100) * 10
	resSum := (nilai - resTax)
	fmt.Println(resTax, resSum)


func ErrorFunc(	data string ) (int, error)  {
	err := strconv.Atoi(data)
	if err != nil {
		return 0, errors.New("Not a Number")
	}else{
		return HitungPajak()
	}
}




