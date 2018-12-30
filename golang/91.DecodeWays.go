func dfs(nums []int, start, end int) int {
	if start > end {
		return 1
	}
	if nums[start] == 0 {
		return 0
	}
	if start == end {
		return 1
	}

	mid := (start + end) / 2

	num := 0
	pre := dfs(nums, start, mid)
	post := dfs(nums, mid+1, end)
	num += pre * post

	if !isValid(nums[mid], nums[mid+1]) {
		return num
	}
	pre = dfs(nums, start, mid-1)
	post = dfs(nums, mid+2, end)
	num += pre * post

	return num
}

func isValid(a, b int) bool {
	value := a*10 + b
	return value >= 10 && value <= 26
}

func numDecodings(s string) int {
	nums := make([]int, len(s))
	for i, c := range s {
		nums[i] = int(c - '0')
	}

	return dfs(nums, 0, len(nums)-1)
}
