func findTheDifference(s string, t string) byte {
	h1 := make([]int, 26)
	h2 := make([]int, 26)
	for _, c := range s {
		h1[c-'a']++
	}
	for _, c := range t {
		h2[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if h2[i] == h1[i] {
			continue
		}
		return byte(i + 'a')
	}
	return 0
}

