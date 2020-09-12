public class Solution {
    public int XorOperation(int n, int start)
    {
        var result = 0;
        for (var i = 0; i < n; i++)
        {
            result ^= start + 2 * i;
        }

        return result;
    }
}

