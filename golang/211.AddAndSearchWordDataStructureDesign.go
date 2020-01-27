const CHILD_NUM = 26

type Trie struct {
	isWord bool
	child  []*Trie
}

type WordDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	dict := WordDictionary{new(Trie)}
	dict.trie.child = make([]*Trie, CHILD_NUM)
	return dict
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	trie := this.trie
	for _, c := range word {
		if trie.child[c-'a'] == nil {
			trie.child[c-'a'] = new(Trie)
			trie.child[c-'a'].child = make([]*Trie, CHILD_NUM)
		}
		trie = trie.child[c-'a']
	}
	trie.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return dfs(this.trie, word)
}

func dfs(node *Trie, word string) bool {
	if node == nil {
		return false
	}
	if len(word) == 0 {
		return node.isWord
	}

	if word[0] != '.' {
		return dfs(node.child[word[0]-'a'], word[1:])
	}

	for i := 0; i < CHILD_NUM; i++ {
		if dfs(node.child[i], word[1:]) {
			return true
		}
	}
	return false
}

