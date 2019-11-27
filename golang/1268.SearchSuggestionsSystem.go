type Node struct {
	word  string
	count int
	child []*Node
}

func newNode() *Node {
	node := new(Node)
	node.child = make([]*Node, 26)
	return node
}

func buildTrie(root *Node, str string) {
	node := root
	for _, c := range str {
		i := int(c - 'a')
		if node.child[i] == nil {
			node.child[i] = newNode()
		}
		node = node.child[i]
	}
	node.word = str
	node.count++
}

func searchTrie(node *Node, result *[]string) {
	if len(node.word) > 0 {
        num := 0
		for len(*result) < 3 && num < node.count {
			*result = append(*result, node.word)
            num++
		}

		if len(*result) == 3 {
			return
		}
	}
	for _, child := range node.child {
		if child == nil {
			continue
		}
		searchTrie(child, result)
		if len(*result) == 3 {
			return
		}
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := newNode()
	for _, product := range products {
		buildTrie(root, product)
	}

	var ans [][]string
	node := root
	for _, c := range searchWord {
		i := int(c - 'a')
		if node.child[i] == nil {
			ans = append(ans, []string{})
			node.child[i] = newNode()
		} else {
			var findings []string
			searchTrie(node.child[i], &findings)
			ans = append(ans, findings)
		}
		node = node.child[i]
	}
	return ans
}

