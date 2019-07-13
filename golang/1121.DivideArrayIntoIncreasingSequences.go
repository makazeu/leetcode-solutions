func canDivideIntoSubsequences(nums []int, K int) bool {
	if len(nums) < K {
		return false
	}
	h := make(map[int]int)
	for _, e := range nums {
		h[e]++
	}
	var counts []int
	for _, v := range h {
		counts = append(counts, v)
	}
	sort.Ints(counts)

	last := -1
	left := 0
	for i, c := range counts {
		if c == last {
			continue
		}
		n := len(counts) - i
		if n >= K {
			left += n - K
		} else {
			left -= K - n
		}
		if left < 0 {
			return false
		}
		last = c
	}
	return true
}
