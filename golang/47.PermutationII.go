func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var answers [][]int

	n := len(nums)
	used := make([]bool, n)
	pos := make([]int, n)
	ans := make([]int, n)
	cur := 0
	pos[0] = -1
	for cur >= 0 {
		if cur == n {
			tmp := make([]int, n)
			copy(tmp, ans)
			answers = append(answers, tmp)
			cur--
			used[pos[cur]] = false
			continue
		}

		var found bool
		for i := pos[cur] + 1; i < n; i++ {
			if used[i] {
				continue
			}
			if pos[cur] != -1 && nums[i] == nums[pos[cur]] {
				continue
			}
			used[i] = true
			ans[cur] = nums[i]
			pos[cur] = i
			cur++
			if cur < n {
				pos[cur] = -1
			}
			found = true
			break
		}

		if !found {
			cur--
			if cur >= 0 {
				used[pos[cur]] = false
			}
		}
	}

	return answers
}
