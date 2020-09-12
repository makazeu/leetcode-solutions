public class Solution
{
    public int NumTeams(int[] rating)
    {
        var N = rating.Length;
        var leftGreater = new int[N];
        var leftLess = new int[N];
        var rightGreater = new int[N];
        var rightLess = new int[N];
        for (var i = 0; i < N; i++)
        {
            var less = 0;
            var greater = 0;
            for (var j = 0; j < i; j++)
            {
                if (rating[j] < rating[i])
                {
                    less++;
                }

                if (rating[j] > rating[i])
                {
                    greater++;
                }
            }

            leftLess[i] = less;
            leftGreater[i] = greater;
            
            less = 0;
            greater = 0;
            for (var j = i+1; j < N; j++)
            {
                if (rating[j] < rating[i])
                {
                    less++;
                }

                if (rating[j] > rating[i])
                {
                    greater++;
                }
            }

            rightLess[i] = less;
            rightGreater[i] = greater;
        }

        var ans = 0;
        for (var i = 0; i < N; i++)
        {
            ans += leftLess[i] * rightGreater[i] + leftGreater[i] * rightLess[i];
        }

        return ans;
    }
}

