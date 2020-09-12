public class Solution
{
    class Item
    {
        public int ch;
        public int count;

        public Item(int ch, int count)
        {
            this.ch = ch;
            this.count = count;
        }
    }

    private int[] count;

    public string LongestDiverseString(int a, int b, int c)
    {
        count = new[] {a, b, c};

        var prevChar = -1;
        var ans = "";

        while (true)
        {
            var next = GetMostCountChar(-1);
            if (count[next] == 0)
            {
                break;
            }
            
            if (count[next] >= 2 && prevChar != next)
            {
                ans += (char)(next + 'a');
                ans += (char)(next + 'a');
                count[next] -= 2;
            }
            else
            {
                ans += (char)(next + 'a');
                count[next]--;
            }

            next = GetMostCountChar(next);
            if (count[next] == 0)
            {
                break;
            }
            ans += (char)(next + 'a');
            count[next]--;
            prevChar = next;
        }

        return ans;
    }

    private int GetMostCountChar(int not)
    {
        var items = new[]
        {
            new Item(0, count[0]),
            new Item(1, count[1]),
            new Item(2, count[2]),
        };
        Array.Sort(items, (x, y) => x.count != y.count
            ? y.count.CompareTo(x.count)
            : x.ch.CompareTo(y.ch));
        return items[0].ch == not ? items[1].ch : items[0].ch;
    }
}

