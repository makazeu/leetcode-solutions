func getNoZeroIntegers(n int) []int {
    for i := 1; i < n; i++ {
		if hasZero(i) || hasZero(n-i) {
			continue
		}
		return []int{i, n-i}
	}
	return nil
}

func hasZero(num int) bool {
	for num > 0 {
		if num % 10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}

