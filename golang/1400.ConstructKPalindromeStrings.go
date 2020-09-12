func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	h := make(map[rune]int)
	for _, c := range s {
		h[c]++
	}
	num := 0
	for _, v := range h {
		num += v % 2
	}
	return num <= k
}

