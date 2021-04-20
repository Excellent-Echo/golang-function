package main

import (
	"fmt"
	"strconv"
)

func bansos() (string, string) {
	var data = map[string]string{
		"name":        "budi",
		"umur":        "45",
		"jarakRumah":  "50",
		"berkeluarga": "ya",
	}

	var keterangan string
	var status bool

	nama, _ := data["name"]
	jarak, _ := strconv.Atoi(data["jarakRumah"])
	berkeluarga, _ := data["berkeluarga"]
	umur, _ := strconv.Atoi(data["umur"])

	if jarak >= 100 && umur <= 20 && berkeluarga == "tidak" {
		status = false
	} else if jarak <= 100 && umur >= 20 && berkeluarga == "ya" {
		status = true
	}

	if status == false {
		keterangan = "tidak layak mendapat bantuan pemerintah"
	} else if status == true {
		keterangan = "layak mendapat bantuan pemerintah"
	}

	return nama, keterangan
}

func main() {
	fmt.Println(bansos())
}
