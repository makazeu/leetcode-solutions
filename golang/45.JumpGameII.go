func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	head, tail := 0, 0
	cur := 0
	step := 0
	fin := len(nums) - 1
	for cur <= tail {
		if cur + nums[cur] >= fin {
			return step + 1
		}
		if cur + nums[cur] > head {
			head = cur + nums[cur]
		}
		if cur == tail {
			step++
			tail = head
		}
		cur++
	}
	return -1
}
