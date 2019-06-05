func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, x := range bills {
		if x == 5 {
			five++
		} else if x == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
				continue
			} else if five >= 3 {
				five -= 3
				continue
			}
			return false
		}
	}
	return true
}

