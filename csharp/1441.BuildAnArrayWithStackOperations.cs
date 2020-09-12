public class Solution {
    public IList<string> BuildArray(int[] target, int n)
    {
        var next = 1;
        var ans = new List<string>();
        foreach (var num in target)
        {
            while (num > next)
            {
                ans.Add("Push");
                ans.Add("Pop");
                next++;
            }
            ans.Add("Push");
            next++;
        }

        return ans;
    }
}

