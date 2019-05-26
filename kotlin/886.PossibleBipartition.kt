class Solution {
    private var label = intArrayOf()
    private var flags = booleanArrayOf()
    private var graph = arrayOf<MutableList<Int>>()
    fun possibleBipartition(N: Int, dislikes: Array<IntArray>): Boolean {
        graph = Array(N) { mutableListOf<Int>() }
        flags = BooleanArray(N)
        dislikes.forEach { graph[it[0] - 1].add(it[1] - 1) }
        for (i in 0 until N) {
            if (flags[i]) continue
            label = IntArray(N) { -1 }
            if (!dfs(i, 0)) return false
        }
        return true
    }

    private fun dfs(k: Int, l: Int): Boolean {
        label[k] = l
        flags[k] = true
        graph[k].forEach {
            if (label[it] == l) return false
            if (label[it] != -1) return@forEach
            if (!dfs(it, l xor 1)) return false
        }
        return true
    }
}

