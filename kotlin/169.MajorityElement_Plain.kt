class Solution {
    fun majorityElement(nums: IntArray): Int {
        var count = 0
        var result = nums[0]
        nums.forEach {
            count += if (it == result) 1 else -1
            if (count == 0) {
                count = 1
                result = it
            }
        }
        return result
    }
}

