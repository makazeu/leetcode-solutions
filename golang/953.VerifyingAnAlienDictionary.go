func isAlienSorted(words []string, order string) bool {
	dict := make(map[rune]rune)
	for i, e := range order {
		dict[e] = rune(i + 'a')
	}

	for i, word := range words {
		var r []rune
		for _, c := range word {
			r = append(r, dict[c])
		}
		words[i] = string(r)

		if i > 0 {
			if words[i] < words[i-1] {
				return false
			}
		}
	}
	return true
}

