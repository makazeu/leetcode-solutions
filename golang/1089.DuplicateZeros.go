func duplicateZeros(arr []int) {
	var nums []int
	for _, e := range arr {
		if e != 0 {
			nums = append(nums, e)
		} else {
			nums = append(nums, 0, 0)
		}
		if len(nums) >= len(arr) {
			break
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = nums[i]
	}
}

