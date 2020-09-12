public class Solution
{
    public bool ThreeConsecutiveOdds(int[] arr)
    {
        for (var i = 0; i <= arr.Length - 3; i++)
        {
            if ((arr[i] & 1) == 0)
            {
                continue;
            }

            var flag = true;
            for (var j = 1; j <= 2; j++)
            {
                if ((arr[i + j] & 1) != 0) continue;
                flag = false;
                i += j;
                break;
            }

            if (flag)
            {
                return true;
            }
        }

        return false;
    }
}

