package function

import "strconv"

func Check(data map[string]string) string {
	jarak, _ := strconv.Atoi(data["jarakRumah"])
	umur, _ := strconv.Atoi(data["umur"])
	berkeluarga := data["berkeluarga"]

	// return jarak, umur, berkeluarga
	if umur > int(20) && jarak < int(100) && berkeluarga == string("ya") {
		return data["name"] + " layak mendapat bantuan dari pemerintah"
	} else {
		return data["name"] + " tidak layak mendapat bantuan dari pemerintah"
	}
}
