public class Solution
{
    public IList<int> FindSmallestSetOfVertices(int n, IList<IList<int>> edges) {
        var result = new List<int>();
        var map = new int[n];
        foreach (var edge in edges)
        {
            map[edge[1]]++;
        }

        for (var i = 0; i < n; i++)
        {
            if (map[i] == 0)
            {
                result.Add(i);
            }
        }

        return result;
    }
}

