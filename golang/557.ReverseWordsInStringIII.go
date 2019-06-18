func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i, str := range strs {
		chars := strings.Split(str, "")
		l := len(chars)
		for j := 0; j < l/2; j++ {
			chars[j], chars[l-j-1] = chars[l-j-1], chars[j]
		}
		strs[i] = ""
		for _, char := range chars {
			strs[i] += char
		}
	}
	return strings.Join(strs, " ")
}

