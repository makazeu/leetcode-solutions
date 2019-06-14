class Solution {
    private var nums = intArrayOf()

    fun subsetsWithDup(nums: IntArray): List<List<Int>> {
        nums.sort()
        this.nums = nums
        val ans = mutableListOf<List<Int>>()
        dfs(mutableListOf(), BooleanArray(nums.size), 0, ans)
        return ans
    }

    private fun dfs(list: MutableList<Int>, used: BooleanArray, pos: Int, ans: MutableList<List<Int>>) {
        if (pos == nums.size) {
            ans.add(list.toList())
            return
        }
        dfs(list, used, pos + 1, ans)

        if (pos > 0 && nums[pos] == nums[pos - 1] && !used[pos - 1]) {
            return
        }
        used[pos] = true
        list.add(nums[pos])
        dfs(list, used, pos + 1, ans)
        list.removeAt(list.size - 1)
        used[pos] = false
    }
}

