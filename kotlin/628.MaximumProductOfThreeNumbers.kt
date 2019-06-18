class Solution {
    fun maximumProduct(nums: IntArray): Int {
        nums.sortDescending()
        var ans = nums[0] * nums[1] * nums[2]
        val len = nums.size
        if (nums[len - 2] < 0) {
            ans = Math.max(ans, nums[0] * nums[len - 1] * nums[len - 2])
        }
        return ans
    }
}

