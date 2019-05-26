let lengthOfLIS = function (nums: number[]): number {
    let dp: number[] = new Array(nums.length);
    for (let i = 0; i < nums.length; i++) {
        dp[i] = 0;
        for (let j = 0; j < i; j++) {
            if (nums[i] > nums[j] && dp[i] < dp[j]) dp[i] = dp[j];
        }
        dp[i]++;
    }
    return dp.length ? Math.max(...dp) : 0;
};

