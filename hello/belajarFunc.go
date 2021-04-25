package main

import (
	"fmt"
	"functionGo/hello"
)

func main() {
	hello.Hello()

	sum, avg, check := hello.Sum(10, 11, 12)

	fmt.Println(sum, avg, check)

	// closure function
	// var getTotal = func(numbers ...int) int {
	// 	var sum int

	// 	for _, num := range numbers {
	// 		sum += num
	// 	}

	// 	return sum
	// }

	// function as parameter
	// result1 := sayHelloFilter("anjing", filterName)
	// result2 := sayHelloFilter("afista", filterName)

	// fmt.Println(result1)
	// fmt.Println(result2)

	// IIFE
	// var check = func(name string, age int) bool {

	// 	if name == "" || age == 0 {
	// 		return true
	// 	} else {
	// 		return false
	// 	}

	// }("afista", 23)

	// fmt.Println(check)
}

// func sayHelloFilter(name string, filter Filt) string { // menerima sebuah fungsi bukan membuat fungsi sendiri
// 	return "hello " + filter(name)
// }

// func filterName(name string) string {
// 	if name == "anjing" || name == "babi" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }
