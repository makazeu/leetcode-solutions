func permute(nums []int) [][]int {
	var answers [][]int
	set := make(map[int]bool)

	n := len(nums)
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
			delete(set, pos[cur])
			continue
		}

		var found bool
		for i := pos[cur] + 1; i < n; i++ {
			if exists := set[i]; exists {
				continue
			}
			set[i] = true
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
				delete(set, pos[cur])
			}
		}
	}

	return answers
}
