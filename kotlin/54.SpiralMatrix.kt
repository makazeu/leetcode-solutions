class Solution {
    private companion object {
        val directions = arrayOf(
            intArrayOf(0, 1),
            intArrayOf(1, 0),
            intArrayOf(0, -1),
            intArrayOf(-1, 0)
        )
    }

    fun spiralOrder(matrix: Array<IntArray>): List<Int> {
        val ans = mutableListOf<Int>()
        if (matrix.isEmpty()) {
            return ans
        }
        val m = matrix.size
        val n = matrix[0].size
        var x = 0
        var y = -1
        var d = 0
        var nx: Int
        var ny: Int
        for (i in 1..m * n) {
            nx = x + directions[d][0]
            ny = y + directions[d][1]

            if (nx < 0 || nx == m || ny < 0 || ny == n || matrix[nx][ny] == Int.MIN_VALUE) {
                d = (d + 1) % 4
                nx = x + directions[d][0]
                ny = y + directions[d][1]
            }
            x = nx
            y = ny
            ans.add(matrix[x][y])
            matrix[x][y] = Int.MIN_VALUE
        }

        return ans
    }
}

