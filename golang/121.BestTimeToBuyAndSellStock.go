func maxProfit(prices []int) int {
	min, ans := math.MaxInt32, 0
	for _, p := range prices {
		if p < min {
			min = p
			continue
		}
		if ans < p-min {
			ans = p - min
		}
	}
	return ans
}

