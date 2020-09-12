public class Solution {
    public IList<bool> KidsWithCandies(int[] candies, int extraCandies)
    {
        var max = candies.Max();
        return candies.Select(i => max - i <= extraCandies).ToArray();
    }
}

