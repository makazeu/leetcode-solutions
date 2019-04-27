class Solution {
    fun minSteps(n: Int): Int {
        if (n == 1) {
            return 0
        }

        val dp = Array(n + 1) { IntArray(n + 1) { n } }
        dp[1][1] = 1

        for (i in 2..n) {
            var min = i
            for (j in 1..i / 2) {
                dp[i][j] = if (dp[i][j] > dp[i - j][j] + 1) dp[i - j][j] + 1 else dp[i][j]
                min = if (min > dp[i][j]) dp[i][j] else min
            }
            dp[i][i] = min + 1
        }

        return dp[n][n] - 1
    }
}

