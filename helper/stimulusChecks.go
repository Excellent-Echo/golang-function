package helper

import (
	"fmt"
	"strconv"
)


func StimulusChecks(data map[string]string) {

	//convert to int
	jarak, _ := strconv.Atoi(data["jarakRumah"])
	umur, _ := strconv.Atoi(data["umur"])

	if jarak < 100 && data["berkeluarga"] == "ya" && umur > 20 {
		fmt.Println(data["name"] + " layak mendapat bantuan dari pemerintah")
	} else {
		fmt.Println(data["name"] + " tidak layak mendapat bantuan dari pemerintah")
	}
}