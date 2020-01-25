func printVertically(s string) []string {
	maxLen := 0
	words := strings.Split(s, " ")
	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	ans := make([][]rune, maxLen)
	for i := range ans {
		ans[i] = make([]rune, len(words))
		for j := range ans[i] {
			ans[i][j] = ' '
		}
	}
	for i, word := range words {
		for j, c := range word {
			ans[j][i] = c
		}
	}

	res := make([]string, maxLen)
	for i := range res {
		res[i] = string(ans[i])
		res[i] = strings.TrimRight(res[i], " ")
	}
	return res
}

