package main

import "fmt"

func main() {
	var (
		max, min int
	)

	array := [5]int{75, 60, 144, 240, 165}
	max = array[0]
	min = array[0]
	for i := 1; i <= len(array)-1; i++ {
		if array[i] > max {
			max = array[i]
		} else if array[i] < min {
			min = array[i]
		}
	}
	if max%2 == 0 {
		fmt.Printf("nilai max = %v, genap", max)
	} else {
		fmt.Printf("nilai max = %v, ganjil", max)
	}
}
