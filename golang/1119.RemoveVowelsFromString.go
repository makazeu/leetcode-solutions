func removeVowels(S string) string {
	var ans []rune
	for _, e := range S {
		if e == 'a' || e == 'e' || e == 'i' || e == 'o' || e == 'u' {
			continue
		}
		ans = append(ans, e)
	}
	return string(ans)
}

