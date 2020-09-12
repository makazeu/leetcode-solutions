public class Solution
{
    public IList<int> MinSubsequence(int[] nums)
    {
        var sum = nums.Sum();
        Array.Sort(nums);
        Array.Reverse(nums);

        var current = 0;
        for (var i = 0; i < nums.Length; i++)
        {
            current += nums[i];
            if (current * 2 <= sum) continue;
            return nums.Take(i + 1).ToList();
        }

        return null;
    }
}

