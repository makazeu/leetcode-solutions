func isArmstrong(N int) bool {
	k := digitNum(N)
	println(k)
	num := N
	sum := int64(0)
	for num > 0 {
		n := num % 10
		num = num / 10
		sum += power(n, k)
		if sum > int64(N) {
			return false
		}
	}
	return sum == int64(N)
}

func digitNum(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func power(n, k int) int64 {
	return int64(math.Pow(float64(n), float64(k)))
}

