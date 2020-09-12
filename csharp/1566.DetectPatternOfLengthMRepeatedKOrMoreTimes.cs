public class Solution {
    public bool ContainsPattern(int[] arr, int m, int k)
    {
        for (var i = 0; i <= arr.Length - m * k; i++)
        {
            var flag = true;
            for (var j = 1; j < k; j++)
            {
                if (Check(arr, i, i + m * j, m)) continue;
                flag = false;
                break;
            }

            if (flag)
            {
                return true;
            }
        }

        return false;
    }

    private bool Check(IReadOnlyList<int> arr, int s1, int s2, int l)
    {
        for (var i = 0; i < l; i++)
        {
            if (arr[s1+i] != arr[s2+i])
            {
                return false;
            }
        }

        return true;
    }
}

