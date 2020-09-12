public class Solution
{
    public int CountGoodTriplets(int[] arr, int a, int b, int c)
    {
        var ans = 0;
        for (var i = 0; i < arr.Length; i++)
        {
            for (var j = i + 1; j < arr.Length; j++)
            {
                if (!Check(arr[i], arr[j], a))
                {
                    continue;
                }

                for (var k = j + 1; k < arr.Length; k++)
                {
                    if (Check(arr[j], arr[k], b) && Check(arr[i], arr[k], c))
                    {
                        ans++;
                    }
                }
            }
        }

        return ans;
    }

    private static bool Check(int a, int b, int delta)
    {
        return Math.Abs(a - b) <= delta;
    }
}

