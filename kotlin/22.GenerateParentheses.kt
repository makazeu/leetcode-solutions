class Solution {
    private val brackets = intArrayOf(0, 1, -1)

    fun generateParenthesis(n: Int): List<String> {
        val results = mutableListOf<String>()
        val current = IntArray(n * 2)
        val num = IntArray(3)
        var pos = 0
        var sum = 0

        while (true) {
            if (pos == n * 2) {
                generateString(results, current)
                pos--
            }

            if (pos == -1) {
                break
            }

            if (current[pos] != 0) {
                sum -= brackets[current[pos]]
                num[current[pos]]--
            }

            var found = false
            for (i in current[pos] + 1..2) {
                if (num[i] == n || sum + brackets[i] < 0) {
                    continue
                }

                current[pos] = i
                sum += brackets[i]
                num[i]++
                found = true
                break
            }

            if (found) {
                pos++
                if (pos < n * 2) {
                    current[pos] = 0
                }
            } else {
                pos--
            }
        }

        return results
    }

    private fun generateString(results: MutableList<String>, current: IntArray) {
        var s = ""
        current.forEach { v ->
            when (v) {
                1 -> s = s.plus("(")
                2 -> s = s.plus(")")
            }
        }
        results.add(s)
    }
}

