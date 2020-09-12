func numberOfSubstrings(s string) int {
	ans := 0
	count := make([]int, 3)
	tail := 0
	for _, c := range s {
		count[c-'a']++
		if count[0] > 0 && count[1] > 0 && count[2] > 0 {
			for count[0] > 0 && count[1] > 0 && count[2] > 0 {
				count[s[tail]-'a']--
				tail++
			}
			ans += tail
			tail--
			count[s[tail]-'a']++
		}
	}
	return ans
}

