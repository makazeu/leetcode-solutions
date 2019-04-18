const (
	even = 0
	odd  = 1
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(key int) {
	node := &Node{key, nil, nil}
	if bst.root == nil {
		bst.root = node
	} else {
		insertNode(bst.root, node)
	}
}

func (bst *BinarySearchTree) PreviousElement(key int) (bool, int) {
	return previousElement(bst.root, key)
}

func (bst *BinarySearchTree) NextElement(key int) (bool, int) {
	return nextElement(bst.root, key)
}

func insertNode(root, node *Node) {
	if node.key < root.key {
		if root.left == nil {
			root.left = node
		} else {
			insertNode(root.left, node)
		}
	} else if node.key > root.key {
		if root.right == nil {
			root.right = node
		} else {
			insertNode(root.right, node)
		}
	}
}

func previousElement(node *Node, key int) (bool, int) {
	if node == nil {
		return false, 0
	}
	if key == node.key {
		return true, node.key
	}
	if key < node.key {
		return previousElement(node.left, key)
	}

	now := node.key
	if f, k := previousElement(node.right, key); f && k > now {
		now = k
	}
	return true, now
}

func nextElement(node *Node, key int) (bool, int) {
	if node == nil {
		return false, 0
	}
	if key == node.key {
		return true, node.key
	}
	if key > node.key {
		return nextElement(node.right, key)
	}

	now := node.key
	if f, k := nextElement(node.left, key); f && k < now {
		now = k
	}
	return true, now
}

func oddEvenJumps(A []int) int {
	N := len(A)
	bst := new(BinarySearchTree)
	posMap := make(map[int]int)
	nextGreater := make([]int, len(A))
	nextLess := make([]int, len(A))

	for i := N - 1; i >= 0; i-- {
		nextLess[i] = -1
		if found, key := bst.PreviousElement(A[i]); found {
			nextLess[i] = posMap[key]
		} else {
			nextLess[i] = -1
		}

		nextGreater[i] = -1
		if found, key := bst.NextElement(A[i]); found {
			nextGreater[i] = posMap[key]
		} else {
			nextGreater[i] = -1
		}

		posMap[A[i]] = i
		bst.Insert(A[i])
	}

	jumpable := make(map[int]map[byte][]int)
	for i := range A {
		jumpable[i] = make(map[byte][]int)
	}
	for i, x := range nextGreater {
		if x == -1 {
			continue
		}
		jumpable[x][even] = append(jumpable[x][even], i)
	}

	for i, x := range nextLess {
		if x == -1 {
			continue
		}
		jumpable[x][odd] = append(jumpable[x][odd], i)
	}

	dp := make([][]bool, len(A))
	for i := range A {
		dp[i] = make([]bool, 2)
	}
	dp[N-1][odd] = true
	dp[N-1][even] = true

	for i := len(A) - 1; i >= 0; i-- {
		if dp[i][odd] {
			for _, pre := range jumpable[i][odd] {
				dp[pre][even] = true
			}
		}
		if dp[i][even] {
			for _, pre := range jumpable[i][even] {
				dp[pre][odd] = true
			}
		}
	}

	ans := 0
	for i := range A {
		if dp[i][odd] {
			ans++
		}
	}
	return ans
}

