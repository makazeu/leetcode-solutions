class Solution {
    fun shipWithinDays(weights: IntArray, D: Int): Int {
        var l = weights.max() ?: 1
        var r = weights.sum()
        var ans = r
        var mid: Int
        while (l <= r) {
            mid = (l + r) / 2
            if (checkOkay(weights, D, mid)) {
                ans = mid
                r = mid - 1
            } else {
                l = mid + 1
            }
        }

        return ans
    }

    private fun checkOkay(weights: IntArray, D: Int, C: Int): Boolean {
        var sum = 0
        var day = 1
        for (weight in weights) {
            if (C < weight) {
                return false
            }
            if (sum + weight <= C) {
                sum += weight
                continue
            }
            day++
            sum = weight

            if (day > D) {
                return false
            }
        }
        return true
    }
}

