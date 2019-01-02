func firstMissingPositive(nums []int) int {
	min, max := math.MaxInt32, math.MinInt32
	for _, n := range nums {
		if n <= 0 {
			continue
		}
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	if min > 1 {
		return 1
	}

	f := make([]bool, max+1)
	for _, n := range nums {
		if n <= 0 {
			continue
		}
		f[n] = true
	}

	for i := 1; i <= max; i++ {
		if f[i] {
			continue
		}
		return i
	}
	return max + 1
}
