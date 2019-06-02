func calc(num, d int) int {
	ans := 0
	power := 1
	for {
		a := num / (power * 10)
		now := a * power

		b := num % (power * 10)
		if b/power == d {
			now += b%power + 1
		} else if b/power > d {
			now += power
		}

		ans += now
		if a == 0 {
			break
		}
		power *= 10
	}
	return ans
}

func calc0(num, d int) int {
	ans := 0
	power := 1
	for {
		a := num / (power * 10)
		now := (a - 1) * power

		b := num % (power * 10)
		if b/power == d {
			now += b%power + 1
		} else if b/power > d {
			now += power
		}

		ans += now
		if a == 0 {
			break
		}
		power *= 10
	}
	return ans
}

func digitsCount(d int, low int, high int) int {
	var a, b int
	if d == 0 {
		a = calc0(high, d)
	} else {
		a = calc(high, d)
	}
	low--
	if low > 0 {
		if d == 0 {
			b = calc0(low, d)
		} else {
			b = calc(low, d)
		}
	}
	return a - b
}

