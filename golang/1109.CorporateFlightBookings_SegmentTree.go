type TreeNode struct {
	num         int
	start, end  int
	left, right *TreeNode
}

func corpFlightBookings(bookings [][]int, n int) []int {
	root := new(TreeNode)
	buildTree(root, 0, n-1)

	for _, e := range bookings {
		addSeg(root, e[0]-1, e[1]-1, e[2])
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = findSeg(root, i)
	}
	return ans
}

func findSeg(node *TreeNode, pos int) int {
	if node.start == node.end {
		return node.num
	}
	res := node.num
	if pos <= (node.start+node.end)/2 {
		res += findSeg(node.left, pos)
	} else {
		res += findSeg(node.right, pos)
	}
	return res
}

func addSeg(node *TreeNode, start, end, num int) {
	if node.start == start && node.end == end {
		node.num += num
		return
	}
	mid := (node.start + node.end) / 2
	if end <= mid {
		addSeg(node.left, start, end, num)
	} else if start > mid {
		addSeg(node.right, start, end, num)
	} else {
		addSeg(node.left, start, mid, num)
		addSeg(node.right, mid+1, end, num)
	}
}

func buildTree(node *TreeNode, left, right int) {
	node.start = left
	node.end = right
	if left == right {
		return
	}
	node.left = new(TreeNode)
	buildTree(node.left, left, (left+right)/2)
	node.right = new(TreeNode)
	buildTree(node.right, (left+right)/2+1, right)
}

