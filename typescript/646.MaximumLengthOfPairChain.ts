let findLongestChain = function(pairs : number[][]) : number {
    pairs.sort((a : number[], b: number[]) : number => {
        return a[0] - b[0];
    });
    let dp : number[] = new Array(pairs.length);
    for (let i = 0; i < pairs.length; i++) {
        dp[i] = 1;
        for (let j = 0; j < i; j++) {
            if (pairs[j][1] >= pairs[i][0]) continue;
            dp[i] = Math.max(dp[i], dp[j] + 1)
        }
    }

    return Math.max(...dp);
};

