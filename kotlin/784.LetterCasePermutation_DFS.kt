class Solution {
    private var S = ""

    fun letterCasePermutation(S: String): List<String> {
        this.S = S
        val ans = mutableListOf<String>()
        dfs("", 0, ans)
        return ans
    }

    private fun dfs(str: String, pos: Int, ans: MutableList<String>) {
        if (str.length == S.length) {
            ans.add(str)
            return
        }
        var s = str
        if (S[pos].isDigit()) {
            s += S[pos]
            dfs(s, pos + 1, ans)
        } else {
            s += S[pos]
            dfs(s, pos + 1, ans)
            s = str + (S[pos].toInt() xor 32).toChar()
            dfs(s, pos + 1, ans)
        }
    }
}

