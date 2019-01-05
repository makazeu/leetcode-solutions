func findRepeatedDnaSequences(s string) []string {
	var answers []string
	occurred := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		str := string(s[i : i+10])
		occur := occurred[str]
		if occur == 1 {
			answers = append(answers, str)
		}
		occurred[str] = occur + 1
	}

	return answers
}
