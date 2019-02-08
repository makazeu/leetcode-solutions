func findPeakElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}

	for i := range nums {
		if checkPeak(nums, i) {
			return i
		}
	}

	return 0
}

func checkPeak(nums []int, pos int) bool {
	if pos == 0 {
		return nums[pos] > nums[pos+1]
	}
	if pos == len(nums)-1 {
		return nums[pos] > nums[pos-1]
	}
	return nums[pos] > nums[pos-1] && nums[pos] > nums[pos+1]
}
