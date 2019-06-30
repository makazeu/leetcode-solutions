func pathInZigZagTree(label int) []int {
	row := getRow(label)
	nums := make([]int, 1)
	now := 0
	var pos int
	for i := 1; i <= row; i++ {
		if i%2 == 1 {
			for j := now + 1; j <= now+(1<<uint(i-1)); j++ {
				nums = append(nums, j)
				if j == label {
					pos = len(nums) - 1
				}
			}
		} else {
			for j := now + (1 << uint(i-1)); j >= now+1; j-- {
				nums = append(nums, j)
				if j == label {
					pos = len(nums) - 1
				}
			}
		}
		now += 1 << uint(i-1)
	}
	ans := make([]int, row)
	for i := row; i >= 1; i-- {
		ans[i-1] = nums[pos]
		pos /= 2
	}
	return ans
}

func getRow(num int) int {
	now := 1
	sum := 0
	for i := 1; ; i++ {
		sum += now
		now *= 2
		if num <= sum {
			return i
		}
	}
}

