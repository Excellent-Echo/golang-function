package hello

import "fmt"

// variadic func
func Total(numbers ...int) {
	// var sum int // temp

	for index, num := range numbers {
		fmt.Println("index :", index, "adalah :", num)
	}
}
