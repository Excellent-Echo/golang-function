package latihan

import "fmt"

func HitungPajak(nilai float32) {
	resTax := (nilai / 100) * 10
	resSum := (nilai - resTax)
	fmt.Println(resTax, resSum)
}
