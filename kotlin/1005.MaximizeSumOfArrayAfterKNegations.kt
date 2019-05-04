class Solution {
    fun largestSumAfterKNegations(A: IntArray, K: Int): Int {
        var k = K
        A.sort()
        for (i in A.indices) {
            if (A[i] >= 0 || k == 0) {
                break
            }
            A[i] = -A[i]
            k--
        }

        var sum = A.sum()
        if (k > 0 && k % 2 == 1) {
            sum -= (A.min() ?: 0) * 2
        }

        return sum
    }
}

