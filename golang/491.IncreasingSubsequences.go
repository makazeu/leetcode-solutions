func findSubsequences(nums []int) [][]int {
	var ans [][]int
	dfs(nums, new([]int), 0, 0, &ans)
	return ans
}

func dfs(nums []int, current *[]int, pos, num int, ans *[][]int) {
	if num > 1 {
		*ans = append(*ans, append([]int(nil), *current...))
	}
	prev := math.MinInt32
	if num > 0 {
		prev = (*current)[num-1]
	}

	m := make(map[int]bool)
	for i := pos; i < len(nums); i++ {
		if prev > nums[i] {
			continue
		}
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		next := append(*current, nums[i])
		dfs(nums, &next, i+1, num+1, ans)
	}
}

