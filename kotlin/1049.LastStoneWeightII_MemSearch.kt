class Solution {
    private var n = 0
    private var stones = intArrayOf()
    private val map = hashMapOf<Triple<Int, Int, Int>, Int>()

    fun lastStoneWeightII(stones: IntArray): Int {
        this.stones = stones
        n = stones.size
        map.clear()
        return dfs(0, 0, n)
    }

    private fun dfs(a: Int, b: Int, left: Int): Int {
        if (left == 0) {
            return Math.abs(a - b)
        }

        val k1 = Triple(a + stones[n - left], b, left - 1)
        val k2 = Triple(a, b + stones[n - left], left - 1)

        return Math.min(
            map.getOrPut(k1, { dfs(a + stones[n - left], b, left - 1) }),
            map.getOrPut(k2, { dfs(a, b + stones[n - left], left - 1) })
        )
    }
}

