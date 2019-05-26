class Solution {
    companion object {
        const val ALPHA_NUM = 26
    }

    private var ans = 0
    private var totalChar = 0
    private val targetChars = ArrayList<Int>()
    private val stickerDict = ArrayList<IntArray>()
    private val map = HashMap<List<Int>, Int>()

    fun minStickers(stickers: Array<String>, target: String): Int {
        ans = Int.MAX_VALUE
        totalChar = 0
        targetChars.clear()
        stickerDict.clear()
        map.clear()

        val targetDict = IntArray(ALPHA_NUM) { -1 }
        target.forEach {
            val k = it - 'a'
            if (targetDict[k] == -1) {
                targetChars.add(0)
                targetDict[k] = totalChar++
            }
            targetChars[targetDict[k]]++
        }


        stickers.forEach {
            var found = false
            val intArray = IntArray(totalChar)
            it.forEach { c ->
                val k = targetDict[c - 'a']
                if (k != -1) {
                    intArray[k]++
                    found = true
                }
            }
            if (found) {
                stickerDict.add(intArray)
            }
        }

        dfs(MutableList(totalChar) { 0 }, 0)

        if (ans == Int.MAX_VALUE) {
            return -1
        }
        return ans
    }

    private fun dfs(current: List<Int>, depth: Int) {
        if (checkEnd(current, depth)) {
            return
        }
        stickerDict.forEach {
            var found = false
            val next = MutableList(totalChar) { i -> current[i] }
            for (i in 0 until totalChar) {
                if (current[i] < targetChars[i] && it[i] > 0) {
                    found = true
                    next[i] = Math.min(targetChars[i], current[i] + it[i])
                }
            }

            if (!found) return@forEach
            if (map.getOrDefault(next, ans) <= depth + 1) return@forEach
            map[next] = depth + 1
            dfs(next, depth + 1)

            if (ans <= depth + 1) return
        }
    }

    private fun checkEnd(current: List<Int>, depth: Int): Boolean {
        for (i in 0 until totalChar) {
            if (current[i] < targetChars[i]) {
                return false
            }
        }
        ans = if (ans > depth) depth else ans
        return true
    }
}

