func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] ^ arr[i]
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		if query[0] == 0 {
			ans[i] = prefix[query[1]]
			continue
		}
		pre := prefix[query[0]-1]
		ans[i] = prefix[query[1]] ^ pre
	}
	return ans
}

