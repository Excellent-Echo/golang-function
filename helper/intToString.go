package helper

import (
	"fmt"
	"strings"
)

func groupOfThreeArray(number int) []int {
	array := []int{}

	for number > 0 {
		array = append(array, number%1000)
		number = number / 1000
	}

	return array
}

func IntToString(input int) {

	satuan := [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	belasan := [...]string{"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}
	puluhan := [...]string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}
	ribuan := [...]string{"", "ribu", "juta", "milyar", "triliun"}

	hasil := []string{}

	if input < 0 {
		hasil = append(hasil, "minus")
		input *= -1
	}

	triplets := groupOfThreeArray(input)


	if input == 0 {
		fmt.Println("nol")
	}

	for i := len(triplets) - 1; i >= 0; i-- {
		triplet := triplets[i]

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		if hundreds == 1 {
			hasil = append(hasil, "seratus")
		} else if hundreds > 0 {
			hasil = append(hasil, satuan[hundreds], "ratus")
		}

		switch tens {
		case 0:
			hasil = append(hasil, satuan[units])
		case 1:
			hasil = append(hasil, belasan[units])
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", puluhan[tens], satuan[units])
				hasil = append(hasil, word)
			} else {
				hasil = append(hasil, puluhan[tens])
			}
		}

		mega := ribuan[i]
		if mega != "" {
			if i == 1 && triplet == 1 {
				hasil = append(hasil[0:len(hasil)-1], "seribu")
			} else {
				hasil = append(hasil, mega)
			}
		}

	}

	fmt.Println(strings.Join(hasil, " ")) 
}