func lengthOfLongestSubstring(s string) int {
	dict := make(map[rune]bool)
	var tail, length int
	chars := []rune(s)

	for _, c := range chars {
		if _, exists := dict[c]; exists {
			length = max(length, len(dict))
			for ; ; tail++ {
				delete(dict, chars[tail])
				if chars[tail] == c {
					tail++
					break
				}
			}
		}
		dict[c] = true
	}

	length = max(length, len(dict))
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
