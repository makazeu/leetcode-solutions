public class Solution
{
    public int BagOfTokensScore(int[] tokens, int P)
    {
        Array.Sort(tokens);
        int i = 0, j = tokens.Length - 1;
        int point = 0, max = 0;

        while (i <= j)
        {
            while (i <= j && P >= tokens[i])
            {
                point++;
                P -= tokens[i++];
            }
            max = Math.Max(max, point);

            if (point == 0)
            {
                break;
            }
            P += tokens[j--];
            point--;
        }
        return max;
    }
}

