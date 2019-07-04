func moveZeroes(nums []int) {
	zeroNum := 0
	pos := 0
	for _, x := range nums {
		if x == 0 {
			zeroNum++
			continue
		}
		nums[pos] = x
		pos++
	}
	for i := 0; i < zeroNum; i++ {
		nums[pos] = 0
		pos++
	}
}

