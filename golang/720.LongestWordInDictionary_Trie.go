type Trie struct {
	word     bool
	children []*Trie
}

const ChildNum = 26

func buildTrie(trie *Trie, word string) {
	for _, c := range word {
		t := c-'a'
		if trie.children[t] == nil {
			trie.children[t] = new(Trie)
			trie.children[t].children = make([]*Trie, ChildNum)
		}
		trie = trie.children[t]
	}
	trie.word = true
}

func longestWord(words []string) string {
	root := new(Trie)
	root.children = make([]*Trie, ChildNum)
	for _, word := range words {
		buildTrie(root, word)
	}

	var ans string
	dfs(root, &ans, "")
	return ans
}

func dfs(trie *Trie, ans *string, current string) {
	if len(current) > len(*ans) || (len(current) == len(*ans) && current < *ans) {
		*ans = current
	}
	for i, child := range trie.children {
		if child == nil {
			continue
		}

		if child.word {
			dfs(child, ans, current + string('a' + i))
		}
	}
}

