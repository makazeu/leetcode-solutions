public class Solution
{
    public int FindLucky(int[] arr)
    {
        var map = new Dictionary<int, int>();
        foreach (var num in arr)
        {
            map[num] = map.GetValueOrDefault(num, 0) + 1;
        }

        return map.Where(pair => pair.Key == pair.Value)
            .Select(pair => pair.Key)
            .Concat(new[] {-1})
            .Max();
    }
}

