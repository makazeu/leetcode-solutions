public class Solution
{
    public int[] Shuffle(int[] nums, int n)
    {
        var ans = new int[n * 2];
        for (var i = 0; i < n * 2; i += 2)
        {
            ans[i] = nums[i / 2];
            ans[i + 1] = nums[n + i / 2];
        }

        return ans;
    }
}

