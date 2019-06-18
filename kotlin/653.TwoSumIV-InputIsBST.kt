/**
 * Example:
 * var ti = TreeNode(5)
 * var v = ti.`val`
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
class Solution {
    private var list = mutableListOf<Int>()
    fun findTarget(root: TreeNode?, k: Int): Boolean {
        list.clear()
        travelTree(root)
        if (list.isEmpty()) return false
        val ans = twoSum(list.toIntArray(), k)
        return ans[0] != 0
    }

    private fun travelTree(node: TreeNode?) {
        if (node == null) return
        travelTree(node.left)
        list.add(node.`val`)
        travelTree(node.right)
    }

    fun twoSum(numbers: IntArray, target: Int): IntArray {
        for (i in 0 until numbers.size) {
            if (numbers[i] * 2 > target) break
            val j = numbers.binarySearch(target - numbers[i], i + 1)
            if (j < 0 || j >= numbers.size || numbers[j] != target - numbers[i]) continue
            return intArrayOf(i + 1, j + 1)
        }
        return intArrayOf(0, 0)
    }
}

