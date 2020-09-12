public class Solution
{
    public int NumSteps(string s)
    {
        var charArray = s.ToCharArray();
        Array.Reverse(charArray);
        var list = charArray.ToList();

        var ans = 0;
        while (list.Count > 1)
        {
            ans++;
            if (list[0] == '0')
            {
                list.RemoveAt(0);
                continue;
            }

            var carry = true;
            for (var i = 0; i < list.Count && carry; i++)
            {
                carry = list[i] == '1';
                list[i] = list[i] == '1' ? '0' : '1';
            }

            if (carry) list.Add('1');
        }

        return ans;
    }
}

