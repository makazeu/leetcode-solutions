func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, c := range magazine {
		m[int(c-'a')]++
	}
	for _, c := range ransomNote {
		m[int(c-'a')]--
		if m[int(c-'a')] < 0 {
			return false
		}
	}
	return true
}

