func reverseOnlyLetters(S string) string {
	var rev []rune
	for i := len(S) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(S[i])) {
			rev = append(rev, rune(S[i]))
		}
	}
	var ans []rune
	pos := 0
	for _, c := range S {
		if unicode.IsLetter(c) {
			ans = append(ans, rev[pos])
			pos++
		} else {
			ans = append(ans, c)
		}
	}
	return string(ans)
}

