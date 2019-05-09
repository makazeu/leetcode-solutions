/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	var ans, str []byte
	dfs(root, 0, &str, &ans)
	return string(ans)
}

func dfs(node *TreeNode, depth int, str, ans *[]byte) {
	depth++
	*str = append([]byte{byte(node.Val + 'a')}, *str...)

	if node.Left == nil && node.Right == nil {
		checkAndUpdateAnswer(*str, ans)
	}

	if node.Left != nil {
		dfs(node.Left, depth, str, ans)
	}
	if node.Right != nil {
		dfs(node.Right, depth, str, ans)
	}

	*str = (*str)[1:]
}

func checkAndUpdateAnswer(str []byte, ans *[]byte) {
	if len(*ans) == 0 {
		*ans = append([]byte(nil), str...)
		return
	}

	if compareLess(str, *ans) {
		*ans = append([]byte(nil), str...)
	}
}

func compareLess(str1, str2 []byte) bool {
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			return str1[i] < str2[i]
		}
	}
	return len(str1) < len(str2)
}

