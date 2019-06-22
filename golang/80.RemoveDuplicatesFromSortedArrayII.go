func removeDuplicates(nums []int) int {
	last := math.MinInt32
	num := 0
	ans := len(nums)
	var newNums []int
	for _, x := range nums {
		if x == last {
			if num == 2 {
				ans--
			} else {
				num++
				newNums = append(newNums, x)
			}
		} else {
			last = x
			num = 1
			newNums = append(newNums, x)
		}
	}
	for i := 0; i < ans; i++ {
		nums[i] = newNums[i]
	}
	return ans
}

