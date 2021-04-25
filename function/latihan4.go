package function

import (
	"errors"
	"strconv"
)

func Check(data map[string]string) (string, error) {
	jarak, err := strconv.Atoi(data["jarakRumah"])
	if err != nil {
		return "0", errors.New("must be number")
	}
	umur, er := strconv.Atoi(data["umur"])
	if er != nil {
		return "0", errors.New("must be number")
	}
	berkeluarga := data["berkeluarga"]

	// return jarak, umur, berkeluarga
	if umur > int(20) && jarak < int(100) && berkeluarga == string("ya") {
		return data["name"] + " layak mendapat bantuan dari pemerintah", err
	} else {
		return data["name"] + " tidak layak mendapat bantuan dari pemerintah", err
	}
}
