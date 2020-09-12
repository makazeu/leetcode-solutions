func countLargestGroup(n int) int {
	h := make(map[int]int)
	for i := 1; i <= n; i++ {
		h[digitSum(i)]++
	}

	maxValue, maxNum := 0, 0
	for _, v := range h {
		if v > maxValue {
			maxValue = v
			maxNum = 1
		} else if v == maxValue {
			maxNum++
		}
	}
	return maxNum
}

func digitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

