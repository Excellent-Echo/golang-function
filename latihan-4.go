package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data = map[string]string{
		"name" : "andi",
		"umur" : "30",
		"jarakRumah" : "50",
		"berkeluarga" : "ya",
	}

	jarakRumah, _ := strconv.Atoi(data["jarakRumah"])
	umur, _ := strconv.Atoi(data["umur"])


	if jarakRumah < 100 && data["berkeluarga"] == "ya" && umur > 20 {
		fmt.Println( data["name"] + " layak mendapat bantuan dari pemerintah")
	}
}

