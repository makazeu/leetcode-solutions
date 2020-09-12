public class Solution
{
    public IList<string> StringMatching(string[] words)
    {
        var ans = new List<string>();
        words = words.OrderBy(x => x.Length).ToArray();
        for (var i = 0; i < words.Length; i++)
        {
            for (var j = words.Length - 1; j >= 0; j--)
            {
                if (i == j) continue;
                if (words[j].Contains(words[i]))
                {
                    ans.Add(words[i]);
                    break;
                }
            }
        }

        return ans;
    }
}

