func generateSentences(synonyms [][]string, text string) []string {
	sets := make(map[string][]string)
	ufs := make(map[string]string)
	var find func(string) string
	find = func(s string) string {
		p := ufs[s]
		if p == s {
			return s
		}
		ufs[s] = find(p)
		return ufs[s]
	}

	merge := func(a, b string) {
		x := find(a)
		y := find(b)
		if x != y {
			ufs[y] = x
		}
	}

	for _, syn := range synonyms {
		ufs[syn[0]] = syn[0]
		ufs[syn[1]] = syn[1]
	}
	for _, syn := range synonyms {
		merge(syn[0], syn[1])
	}
	for k := range ufs {
		p := find(k)
		sets[p] = append(sets[p], k)
	}

	var sentences []string
	words := strings.Split(text, " ")
	l := len(words)
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == l {
			sentence := strings.Join(words, " ")
			sentences = append(sentences, sentence)
			return
		}
		set, exists := sets[find(words[pos])]
		if !exists {
			dfs(pos + 1)
			return
		}
		for _, s := range set {
			words[pos] = s
			dfs(pos + 1)
		}
	}

	dfs(0)
	sort.Strings(sentences)
	return sentences
}

