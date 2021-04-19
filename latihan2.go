package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//Input data
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("masukkan input : ")
	scanner.Scan()
	nominal, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("IDR %v,00", nominal)
}
