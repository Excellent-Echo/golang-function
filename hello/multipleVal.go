package hello

// return multiple value
func Sum(x, y, z int) (int, int, bool) { // 100 return juga gpp
	if x <= 0 || y <= 0 || z <= 0 {
		return 0, 0, true
	} else {
		sum := x + y + z
		avg := sum / 3
		return sum, avg, false
	}
}
