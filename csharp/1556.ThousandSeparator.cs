public class Solution {
    public string ThousandSeparator(int n) {
        if (n == 0)
        {
            return "0";
        }

        var result = "";
        while (n > 0)
        {
            if (result.Length % 4 == 3)
            {
                result = $".{result}";
            }
            
            result = $"{n % 10}{result}";
            n /= 10;
        }

        return result;
    }
}

