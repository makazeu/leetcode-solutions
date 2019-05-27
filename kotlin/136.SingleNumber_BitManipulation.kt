class Solution {
    fun singleNumber(nums: IntArray): Int {
        return nums.fold(0) { acc, i -> acc xor i }
    }
}

