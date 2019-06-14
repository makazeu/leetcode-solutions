class Solution {
    private var nums = intArrayOf()

    fun subsets(nums: IntArray): List<List<Int>> {
        this.nums = nums
        val ans = mutableListOf<List<Int>>()
        dfs(mutableListOf(), 0, ans)
        return ans
    }

    private fun dfs(list: MutableList<Int>, pos: Int, ans: MutableList<List<Int>>) {
        if (pos == nums.size) {
            ans.add(list.toList())
            return
        }
        dfs(list, pos + 1, ans)
        list.add(nums[pos])
        dfs(list, pos + 1, ans)
        list.removeAt(list.size - 1)
    }
}

