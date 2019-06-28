func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 0
	for d := 1; d*d <= num; d++ {
		if num%d != 0 {
			continue
		}
		sum += d
		v := num / d
		if v == num || v == d {
			continue
		}
		sum += v
	}
	return sum == num
}

