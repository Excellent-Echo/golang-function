package main

import "fmt"

func GenapGanjil(number ...int) {
	var num []int
	var check string
	var max, min int
	num = (number)

	// array := [5]int{75, 60, 144, 241, 165}
	max = num[0]
	min = num[0]

	//max min
	if len(number) > 0 {
		for i := 1; i < len(num); i++ {

			if num[i] > max {
				max = num[i]
			} else if num[i] < min {
				min = num[i]
			}
			if max%2 == 0 {
				check = "angka terbanyak adalah genap"
			} else if max%2 != 0 {
				check = "angka terbanyak adalah ganjil"
			}
		}
	} else if len(number) == 0 {
		check = "tidak ada angka"
	}
	fmt.Println(check)
}

func main() {

	GenapGanjil(1, 2, 3, 4, 5)
	GenapGanjil(4, 2)
	GenapGanjil(10, 20, 30, 13)
	GenapGanjil(30, 13, 13, 77, 33, 55, 17, 13)

}
