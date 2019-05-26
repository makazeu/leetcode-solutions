let findNumberOfLIS = function (nums: number[]): number {
    let dp: number[] = new Array(nums.length);
    let count: number[] = new Array(nums.length);

    let max = 0, ans = 0;
    for (let i = 0; i < nums.length; i++) {
        dp[i] = 0;
        count[i] = 1;
        for (let j = 0; j < i; j++) {
            if (nums[j] >= nums[i]) continue;
            if (dp[j] < dp[i]) continue;
            if (dp[j] == dp[i]) {
                count[i] += count[j];
            } else {
                dp[i] = dp[j];
                count[i] = count[j];
            }
        }
        dp[i]++;
        if (dp[i] > max) {
            max = dp[i];
            ans = count[i];
        } else if (dp[i] == max) {
            ans += count[i];
        }
    }
    return ans;
};

