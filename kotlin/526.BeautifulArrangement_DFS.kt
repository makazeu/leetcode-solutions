class Solution {
    private var N = 0
    private var ans = 0
    private var used = booleanArrayOf()

    fun countArrangement(N: Int): Int {
        this.N = N
        ans = 0
        used = BooleanArray(N + 1)
        dfs(1)
        return ans
    }

    private fun dfs(pos: Int) {
        if (pos == N + 1) {
            ans++
            return
        }
        for (i in 1..N) {
            if (used[i]) continue
            if (i % pos != 0 && pos % i != 0) continue
            used[i] = true
            dfs(pos + 1)
            used[i] = false
        }
    }
}

