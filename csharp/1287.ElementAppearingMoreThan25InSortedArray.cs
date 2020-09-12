public class Solution {
    public int FindSpecialInteger(int[] arr)
    {
        var n = arr.Length / 4 + 1;
        for (int i = 0; i <= arr.Length - n; i++)
        {
            if (arr[i] == arr[i+n-1])
            {
                return arr[i];
            }
        }

        return -1;
    }
}

