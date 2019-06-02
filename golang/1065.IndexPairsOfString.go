func check(text, word string, start int) bool {
	if start+len(word) > len(text) {
		return false
	}
	for i := range word {
		if word[i] != text[start] {
			return false
		}
		start++
	}
	return true
}

type Words []string

func (words Words) Len() int {
	return len(words)
}

func (words Words) Less(i, j int) bool {
	return len(words[i]) < len(words[j])
}

func (words Words) Swap(i, j int) {
	words[i], words[j] = words[j], words[i]
}

func indexPairs(text string, words []string) [][]int {
	sort.Sort(Words(words))
	var ans [][]int
	for i := range text {
		for _, word := range words {
			if check(text, word, i) {
				ans = append(ans, []int{i, i + len(word) - 1})
			}
		}
	}
	return ans
}

