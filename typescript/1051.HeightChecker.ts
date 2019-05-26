let heightChecker = function (heights: number[]): number {
    let sortedHeights = [...heights];
    sortedHeights.sort((a, b) => a - b);
    let ans = 0;
    for (let i = 0; i < heights.length; i++) {
        ans += heights[i] != sortedHeights[i] ? 1 : 0;
    }
    return ans;
};

