const ChildNum = 26

type Node struct {
    isWord bool
    child []*Node
}

type Trie struct {
    root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
    trie := Trie{}
    trie.root = new(Node)
    trie.root.child = make([]*Node, ChildNum)
    return trie
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    var t int
    node := this.root
    for _, c := range word {
        t = int(c-'a')
        if node.child[t] == nil {
            node.child[t] = new(Node)
            node.child[t].child = make([]*Node, ChildNum)
        }
        node = node.child[t]
    }
    node.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    var t int
    node := this.root
    for _, c := range word {
        t = int(c-'a')
        if node.child[t] == nil {
            return false
        }
        node = node.child[t]
    }
    return node.isWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    var t int
    node := this.root
    for _, c := range prefix {
        t = int(c-'a')
        if node.child[t] == nil {
            return false
        }
        node = node.child[t]
    }
    return searchAny(node)
}

func searchAny(node *Node) bool {
    if node.isWord {
        return true
    }
    for i := 0; i < ChildNum; i++ {
        if node.child[i] == nil {
            continue
        }
        if searchAny(node.child[i]) {
            return true
        }
    }
    return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

