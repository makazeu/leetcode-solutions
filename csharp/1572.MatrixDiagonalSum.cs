public class Solution
{
    public int DiagonalSum(int[][] mat)
    {
        var n = mat.Length;
        var sum = 0;
        for (var i = 0; i < n; i++)
        {
            sum += mat[i][i] + mat[i][n - i - 1];
        }

        if ((n & 1) == 1)
        {
            sum -= mat[n / 2][n / 2];
        }

        return sum;
    }
}

