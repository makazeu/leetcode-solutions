class Solution {
    fun maxUncrossedLines(A: IntArray, B: IntArray): Int {
        val dp = Array(A.size + 1) { IntArray(B.size + 1) { 0 } }

        for (i in 1..A.size) {
            for (j in 1..B.size) {
                dp[i][j] = maxOf(
                    dp[i][j - 1], dp[i - 1][j],
                    if (A[i - 1] == B[j - 1]) dp[i - 1][j - 1] + 1 else dp[i - 1][j - 1]
                )
            }
        }
        return dp[A.size][B.size]
    }
}

