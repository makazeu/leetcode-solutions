class Solution {
    fun prefixesDivBy5(A: IntArray): BooleanArray {
        val booleanArray = BooleanArray(A.size)
        var num = 0
        for (i in A.indices) {
            num = (num * 2 + A[i]) % 5
            booleanArray[i] = num == 0
        }
        return booleanArray
    }
}

