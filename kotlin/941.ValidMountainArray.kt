class Solution {
    fun validMountainArray(A: IntArray): Boolean {
        if (A.size < 3) {
            return false
        }

        var found = 0
        for (i in 1 until A.size) {
            if (A[i] == A[i - 1]) return false
            if (A[i] < A[i - 1]) {
                found = i - 1
                break
            }
        }
        if (found == 0) return false
        for (i in found + 1 until A.size) {
            if (A[i] >= A[i - 1]) return false
        }
        return true
    }
}

