public class Solution
{
    public string[] GetFolderNames(string[] names)
    {
        var lastIndexes = new Dictionary<string, int>();
        var prevNames = new HashSet<string>();

        for (var i = 0; i < names.Length; i++)
        {
            var name = names[i];
            if (!prevNames.Contains(name))
            {
                prevNames.Add(name);
                lastIndexes[name] = 0;
                continue;
            }

            var nextIndex = (lastIndexes.ContainsKey(name) ? lastIndexes[name] : 0) + 1;
            var nextName = $"{name}({nextIndex})";
            while (prevNames.Contains(nextName))
            {
                nextIndex++;
                nextName = $"{name}({nextIndex})";
            }

            names[i] = nextName;
            lastIndexes[name] = nextIndex;
            prevNames.Add(nextName);
        }

        return names;
    }
}

