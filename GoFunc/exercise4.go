package GoFunc

import (
	"fmt"
	"strconv"
)

func Relief(data map[string]string) {
	houseDistance, _ := strconv.Atoi(data["jarakRumah"])
	age, _ := strconv.Atoi(data["umur"])

	switch houseDistance < 100 && data["berkeluarga"] == "ya" && age > 20 {
	case true:
		fmt.Println(data["name"] + " layak mendapat bantuan dari pemerintah")
	case false:
		fmt.Println(data["name"] + " tidak layak mendapat bantuan dari pemerintah")
	}
}