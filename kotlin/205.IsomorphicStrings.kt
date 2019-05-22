class Solution {
    fun isIsomorphic(s: String, t: String): Boolean {
        val pmap = HashMap<Int, Int>()
        val rmap = HashMap<Int, Int>()
        for (i in s.indices) {
            val a = s[i] - 'a'
            val b = t[i] - 'a'
            if (pmap[a] == null && rmap[b] == null) {
                pmap[a] = b
                rmap[b] = a
            } else if (pmap[a] != b || rmap[b] != a) {
                return false
            }
        }
        return true
    }
}

