package main

import (
	"fmt"
	"strconv"
)

func main() {

	var data = []map[string]string{
		{"name": "andi", "umur": "30", "jarakRumah": "50", "berkeluarga": "ya"},
		{"name": "santi", "umur": "19", "jarakRumah": "80", "berkeluarga": "ya"},
		{"name": "budi", "umur": "45", "jarakRumah": "120", "berkeluarga": "ya"},
	}
	for _, i := range data {
		nama, _ := i["name"]
		jarak := strconv.Atoi(i["jarakRumah"])
		age := strconv.Atoi(i["umur"])
		keluarga := i["berkeluarga"]

		if jarak < 100 || age > 20 || keluarga == "ya" {
			fmt.Printf("%v layak mendapatkan bantuan dari pemerintah\n", nama)
		} else {
			fmt.Printf("%v tidak layak mendapatkan bantuan dari pemerintah\n", nama)
		}

	}
}
