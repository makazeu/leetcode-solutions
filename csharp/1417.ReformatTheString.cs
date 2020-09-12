public class Solution
{
    public string Reformat(string s)
    {
        var list = new[]{new List<char>(), new List<char>() };
        foreach (var c in s)
        {
            list[char.IsNumber(c) ? 0 : 1].Add(c);
        }

        if (Math.Abs(list[0].Count - list[1].Count) >= 2)
        {
            return string.Empty;
        }

        var index = 0;
        if (list[1].Count > list[0].Count)
        {
            index = 1;
        }

        var ans = "";
        var p = new[] {0, 0};
        while (p[0] < list[0].Count || p[1] < list[1].Count)
        {
            ans += list[index][p[index]];
            p[index]++;
            index ^= 1;
        }

        return ans;
    }
}

