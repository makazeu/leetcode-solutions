public class Solution {
    public int MinStartValue(int[] nums) {
        var min = nums[0];
        var sum = 0;
        foreach(var num in nums) {
            sum += num;
            if (sum < min) {
                min = sum;
            }
        }
        return min > 0 ? 1 : 1 - min;
    }
}

