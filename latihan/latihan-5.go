package latihan

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func groupOfThreeArray(number int) []int {
	array := []int{}

	for number > 0 {
		array = append(array, number%1000)
		number = number / 1000
		// fmt.Printf("%v %v\n", array, number)
	}

	return array
}

func ChangeNumtoStr(input string) (string, error) {
	number, err := strconv.Atoi(input)

	satuan := [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	belasan := [...]string{"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}
	puluhan := [...]string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}
	ribuan := [...]string{"", "ribu", "juta", "milyar", "triliun"}

	hasil := []string{}

	if number < 0 {
		hasil = append(hasil, "minus")
		number *= -1
	}

	// ubah numberan ke array triplets
	triplets := groupOfThreeArray(number)

	// error handling
	if err != nil {
		return "", errors.New("Input incorrect: Please input a number")
	}

	if input == "0" {
		return "nol", nil
	}

	for i := len(triplets) - 1; i >= 0; i-- {
		triplet := triplets[i]
		// fmt.Printf("Triplet: %d (i=%d)\n", triplet, i)

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		// fmt.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)

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
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", puluhan[tens], satuan[units])
				hasil = append(hasil, word)
			} else {
				hasil = append(hasil, puluhan[tens])
			}
			break
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

	return strings.Join(hasil, " "), nil
}
