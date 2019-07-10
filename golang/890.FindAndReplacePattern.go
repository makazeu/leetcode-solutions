func findAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	for _, word := range words {
		if len(word) != len(pattern) {
			continue
		}
		a := make([]int, 26)
		b := make([]int, 26)
		flag := true
		for i, c := range word {
			if a[c-'a'] != 0 && a[c-'a'] != int(pattern[i]) {
				flag = false
				break
			}
			if b[pattern[i]-'a'] != 0 && b[pattern[i]-'a'] != int(c) {
				flag = false
				break
			}
			a[c-'a'] = int(pattern[i])
			b[pattern[i]-'a'] = int(c)
		}
		if flag {
			ans = append(ans, word)
		}
	}
	return ans
}

