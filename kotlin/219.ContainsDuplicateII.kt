class Solution {
    fun containsNearbyDuplicate(nums: IntArray, k: Int): Boolean {
        val map = HashMap<Int, MutableList<Int>>()
        for (i in nums.indices) {
            map.getOrPut(nums[i], { mutableListOf() }).add(i)
        }
        map.values.filter { it.size > 1 }
            .forEach {
                it.sort()
                for (i in 1 until it.size) {
                    if (it[i] - it[i - 1] <= k) {
                        return true
                    }
                }
            }
        return false
    }
}

