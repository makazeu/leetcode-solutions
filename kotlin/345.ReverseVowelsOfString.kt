class Solution {
    private val vowels = charArrayOf('a', 'e', 'i', 'o', 'u')

    fun reverseVowels(s: String): String {
        val array = s.toCharArray()
        var l = 0
        var r = array.size - 1
        while (l < r) {
            while (l < r && checkNotVowel(array[l])) l++
            while (l < r && checkNotVowel(array[r])) r--
            if (l >= r) break

            array[l] = array[r].also { array[r] = array[l] }
            l++
            r--
        }
        return String(array)
    }

    private fun checkNotVowel(c: Char): Boolean {
        return c.toLowerCase() !in vowels
    }
}

