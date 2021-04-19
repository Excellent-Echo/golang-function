package helper

import "strconv"

func BantuanCheck(data map[string]string) (res string) {
	umur, _ := strconv.Atoi(data["umur"])
	jarak, _ := strconv.Atoi(data["jarakRumah"])
	berkeluarga := data["berkeluarga"]

	if umur > 20 && jarak < 100 && berkeluarga == "ya" {
		res = data["name"] + " layak mendapatkan bantuan dari pemerintah"
	} else {
		res = data["name"] + " tidak layak mendapat bantuan dari pemerintah"
	}
	return
}
