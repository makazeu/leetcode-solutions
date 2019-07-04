func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	sum := 0
	for _, e := range nums {
		sum += e
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

