class Solution {
    fun pivotIndex(nums: IntArray): Int {
        if (nums.isEmpty()) return -1
        val n = nums.size
        val sums = IntArray(n)
        sums[0] = nums[0]
        for (i in 1 until n) {
            sums[i] = sums[i - 1] + nums[i]
        }

        var left: Int
        var right: Int
        for (i in nums.indices) {
            left = if (i - 1 >= 0) sums[i - 1] else 0
            right = if (i + 1 < n) sums[n - 1] - sums[i] else 0
            if (left == right) {
                return i
            }
        }
        return -1
    }
}

