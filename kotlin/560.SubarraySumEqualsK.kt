class Solution {
    fun subarraySum(nums: IntArray, k: Int): Int {
        val map = HashMap<Int, Int>()
        val sums = IntArray(nums.size + 1)
        for (i in 1..nums.size) {
            sums[i] = sums[i - 1] + nums[i - 1]
            map[sums[i]] = map.getOrDefault(sums[i], 0) + 1
        }

        var ans = 0
        for (i in 0 until nums.size) {
            if (i > 0) {
                map[sums[i]] = map[sums[i]]!! - 1
            }
            ans += map.getOrDefault(k + sums[i], 0)
        }

        return ans
    }
}

