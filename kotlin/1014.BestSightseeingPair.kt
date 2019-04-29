class Solution {
    fun maxScoreSightseeingPair(A: IntArray): Int {
        val N = A.size
        val L = IntArray(N)
        val R = IntArray(N)

        L[0] = A[0]
        for (i in 1..N - 2) {
            L[i] = if (L[i - 1] > A[i] + i) L[i - 1] else A[i] + i
        }

        R[N - 1] = A[N - 1] - N + 1
        for (j in N - 2 downTo 1) {
            R[j] = if (R[j + 1] > A[j] - j) R[j + 1] else A[j] - j
        }

        var ans = 0
        for (i in 0..N - 2) {
            ans = if (L[i] + R[i + 1] > ans) L[i] + R[i + 1] else ans
        }
        return ans
    }
}

