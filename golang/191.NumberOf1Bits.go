func hammingWeight(num uint32) int {
	ans := 0
	for i := uint(0); i < 32; i++ {
		if 1<<i&num > 0 {
			ans++
		}
	}
	return ans
}

