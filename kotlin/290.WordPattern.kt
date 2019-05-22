class Solution {
    fun wordPattern(pattern: String, str: String): Boolean {
        val s = str.split(" ")
        if (s.size != pattern.length) return false
        val p = arrayOfNulls<String>(26)
        val map = HashMap<String, Int>()
        for (i in s.indices) {
            val c = pattern[i] - 'a'
            if (p[c] == null) p[c] = s[i]
            map.putIfAbsent(s[i], c)

            if (p[c] != s[i]) return false
            if (map[s[i]] != c) return false
        }
        return true
    }
}

