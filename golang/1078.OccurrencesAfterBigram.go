func findOcurrences(text string, first string, second string) []string {
	str := strings.Split(text, " ")
	var ans []string
	for i := 0; i < len(str)-2; i++ {
		if strings.EqualFold(str[i], first) && strings.EqualFold(str[i+1], second) {
			ans = append(ans, str[i+2])
		}
	}
	return ans
}

