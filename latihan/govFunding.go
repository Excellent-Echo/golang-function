package latihan

import (
	"fmt"
	"strconv"
)

func GovFunding(data map[string]string) string {
	umur, _ := strconv.Atoi(data["umur"])
	jarakRumah, _ := strconv.Atoi(data["jarakRumah"])

	if jarakRumah < 100 && umur > 20 && data["berkeluarga"] == "ya" {
		return fmt.Sprintf("%s layak mendapatkan bantuan", data["name"])
	} else {
		return fmt.Sprintf("%s tidak layak mendapatkan bantuan", data["name"])
	}
}
