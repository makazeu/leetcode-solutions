class Solution {
    private companion object {
        val directions = arrayOf(
            intArrayOf(0, 1),
            intArrayOf(1, 0),
            intArrayOf(0, -1),
            intArrayOf(-1, 0)
        )
    }

    fun generateMatrix(n: Int): Array<IntArray> {
        val mat = Array(n) { IntArray(n) }
        var x = 0
        var y = -1
        var d = 0
        var nx: Int
        var ny: Int
        for (i in 1..n * n) {
            nx = x + directions[d][0]
            ny = y + directions[d][1]

            if (nx < 0 || nx == n || ny < 0 || ny == n || mat[nx][ny] != 0) {
                d = (d + 1) % 4
                nx = x + directions[d][0]
                ny = y + directions[d][1]
            }
            x = nx
            y = ny
            mat[x][y] = i
        }
        
        return mat
    }
}

