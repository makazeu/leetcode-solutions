public class Solution {
    public string DestCity(IList<IList<string>> paths) {
        var incoming = new HashSet<string>();
        var outgoing = new HashSet<string>();
        foreach (var path in paths)
        {
            incoming.Add(path[1]);
            outgoing.Add(path[0]);
        }

        return incoming.Single(city => !outgoing.Contains(city));
    }
}

