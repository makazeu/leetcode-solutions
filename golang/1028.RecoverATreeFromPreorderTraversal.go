/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	var depth, value []int
	flag := false
	d, v := 0, 0
	for _, c := range S {
		if c >= '0' && c <= '9' {
			if flag {
				v = v*10 + int(c-'0')
			} else {
				flag = true
				v = int(c - '0')
			}
		} else {
			if flag {
				depth = append(depth, d)
				value = append(value, v)
				d = 1
				flag = false
			} else {
				d++
			}
		}
	}
	depth = append(depth, d)
	value = append(value, v)

	fmt.Println(depth)
	fmt.Println(value)

	root := new(TreeNode)
	recoverTree(root, depth, value, 0, 0)
	return root
}

func recoverTree(root *TreeNode, depth, value []int, dep, pos int) int {
	if dep == depth[pos] {
		root.Val = value[pos]
		pos++
	}

	if pos >= len(value) {
		return pos
	}

	if depth[pos] > dep {
		if root.Left == nil {
			root.Left = new(TreeNode)
			pos = recoverTree(root.Left, depth, value, dep + 1, pos)
		}
	}

	if pos >= len(value) {
		return pos
	}

	if depth[pos] > dep {
		if root.Right == nil {
			root.Right = new(TreeNode)
			pos = recoverTree(root.Right, depth, value, dep + 1, pos)
		}
	}

	return pos
}

