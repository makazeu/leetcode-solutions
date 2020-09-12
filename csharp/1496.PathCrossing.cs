public class Solution
{
    public bool IsPathCrossing(string path)
    {
        var set = new HashSet<KeyValuePair<int, int>>();
        int x = 0, y = 0;
        set.Add(new KeyValuePair<int, int>(x, y));
        foreach (var c in path)
        {
            switch (c)
            {
                case 'N':
                    y++;
                    break;
                case 'E':
                    x++;
                    break;
                case 'S':
                    y--;
                    break;
                case 'W':
                    x--;
                    break;
            }

            var pair = new KeyValuePair<int, int>(x, y);
            if (set.Contains(pair))
            {
                return true;
            }

            set.Add(pair);
        }

        return false;
    }
}

