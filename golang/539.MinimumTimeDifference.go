func findMinDifference(timePoints []string) int {
	nums := make([]int, len(timePoints))
	for i, t := range timePoints {
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		nums[i] = hour*60 + minute
	}
	sort.Ints(nums)
	min := nums[0] + 1440 - nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] < min {
			min = nums[i+1] - nums[i]
		}
	}
	return min
}

