func isSolvable(words []string, result string) bool {
	charsMap := make(map[rune]bool)
	zeroMap := make([]bool, 26)
	for _, word := range words {
		zeroMap[word[0]-'A'] = true
		for _, c := range word {
			charsMap[c] = true
		}
	}
	zeroMap[result[0]-'A'] = true
	for _, c := range result {
		charsMap[c] = true
	}

	var chars []rune
	for c := range charsMap {
		chars = append(chars, c)
	}

	charValue := make([]int, 26)
	used := make([]bool, 10)
	return dfs(words, result, chars, 0, charValue, used, zeroMap)
}

func dfs(words []string, result string, chars []rune, pos int, charValue []int, used, zeroMap []bool) bool {
	if pos == len(chars) {
		return check(words, result, charValue)
	}
	for i := 0; i <= 9; i++ {
		if used[i] {
			continue
		}
		if i == 0 && zeroMap[chars[pos]-'A'] {
			continue
		}
		charValue[chars[pos]-'A'] = i
		used[i] = true

		if dfs(words, result, chars, pos+1, charValue, used, zeroMap) {
			return true
		}
		used[i] = false
	}
	return false
}

func check(words []string, result string, charValue []int) bool {
	sum := 0
	for _, word := range words {
		now := 0
		for _, c := range word {
			now = now*10 + charValue[c-'A']
		}
		sum += now
	}

	res := 0
	for _, c := range result {
		res = res*10 + charValue[c-'A']
	}

	return sum == res
}

