class Solution {
    fun numSubarrayProductLessThanK(nums: IntArray, k: Int): Int {
        if (k <= 1) {
            return 0
        }
        var ans = 0
        var prod = 1
        var i = 0
        var j = 0

        while (j < nums.size) {
            prod *= nums[j]
            while (i <= j + 1 && prod >= k && i < nums.size) {
                prod /= nums[i]
                i++
            }
            ans += (j - i + 1)
            j++
        }

        return ans
    }
}

