func canJump(nums []int) bool {
	max := 0
	target := len(nums) - 1
	for i, num := range nums {
		if i > max {
			return false
		}
		max = Max(max, i+num)

		if max >= target {
			return true
		}
	}
	return max >= target
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
