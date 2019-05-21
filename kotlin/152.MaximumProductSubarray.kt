class Solution {
    fun maxProduct(nums: IntArray): Int {
        if (nums.size == 1) {
            return nums[0]
        }
        
        var p = 0
        var n = 0
        var ans = 0

        var np : Int

        for (num in nums) {
            when {
                num > 0 -> {
                    p = Math.max(p * num, num)
                    n = Math.min(n * num, num)
                }
                num < 0 -> {
                    np = Math.max(n * num, 0)
                    n = Math.min(p * num, num)
                    p = np
                }
                else -> {
                    p = 0
                    n = 0
                }
            }

            ans = Math.max(ans, p)
        }

        return ans
    }
}

