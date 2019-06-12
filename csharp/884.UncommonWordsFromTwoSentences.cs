public class Solution {
    public string[] UncommonFromSentences(string A, string B) {
        var dict1 = A.Split(" ")
            .GroupBy(s => s)
            .Select(s => new {Str = s.Key, Count = s.Count()})
            .ToDictionary(g => g.Str, g => g.Count);

        var dict2 = B.Split(" ")
            .GroupBy(s => s)
            .Select(s => new {Str = s.Key, Count = s.Count()})
            .ToDictionary(g => g.Str, g => g.Count);

        return dict1.Keys.ToList().FindAll(e => dict1[e] == 1)
            .Union(dict2.Keys.ToList().FindAll(e => dict2[e] == 1))
            .ToList()
            .FindAll(str => dict1.ContainsKey(str) ^ dict2.ContainsKey(str))
            .ToArray();
    }
}

