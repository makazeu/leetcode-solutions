class Solution {
    fun hasGroupsSizeX(deck: IntArray): Boolean {
        if (deck.size < 2) return false
        val map = HashMap<Int, Int>()
        deck.forEach { map[it] = map.getOrDefault(it, 0) + 1 }
        return map.values.reduce { acc, i -> gcd(acc, i) } >= 2
    }

    private fun gcd(a: Int, b: Int): Int = if (b == 0) a else gcd(b, a % b)
}

