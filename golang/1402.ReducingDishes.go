func maxSatisfaction(satisfaction []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	n := len(satisfaction)
	ans := 0
	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j >=1; j-- {
			sum += j * satisfaction[i-j]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

