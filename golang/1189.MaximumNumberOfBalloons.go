func maxNumberOfBalloons(text string) int {
	count := make([]int, 26)
	for _, c := range text {
		count[c-'a']++
	}

	ans := count[1]
	ans = min(ans, count[0])
	ans = min(ans, count['l'-'a']/2)
	ans = min(ans, count['o'-'a']/2)
	ans = min(ans, count['n'-'a'])
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

